package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"golang.org/x/sync/errgroup"

	"github.com/minghsu0107/gocrypt/pkg/api"
	"github.com/minghsu0107/gocrypt/pkg/display/allcoin"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gocrypt",
	Short: "A terminal application to watch crypto prices!",
	Long:  `gocrypt is a TUI based application that monitors cryptocurrency prices in real time, written in Go.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		// Context and errgroup used to manage routines
		eg, ctx := errgroup.WithContext(context.Background())
		dataChannel := make(chan api.AssetData)

		// Flag to determine if data must be sent when viewing per coin prices
		sendData := true

		// Fetch Coin Assets
		eg.Go(func() error {
			return api.GetAssets(ctx, dataChannel, &sendData)
		})

		// Display UI for overall coins
		eg.Go(func() error {
			return allcoin.DisplayAllCoins(ctx, dataChannel, &sendData)
		})

		if err := eg.Wait(); err != nil {
			if err.Error() != "UI Closed" {
				return err
			}
		}
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.gocrypt.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	setDefault()

	viper.SetConfigType("yaml")
	viper.SetEnvPrefix("GOCRYPT")
	viper.AutomaticEnv()                                   // read in environment variables that match
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // eg. GOCRYPT_CURRENCY_INITUNIT=new-taiwan-dollar ./gocrypt

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".gocrypt" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".gocrypt")
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

func setDefault() {
	viper.SetDefault("currency.initUnit", "united-states-dollar")
}
