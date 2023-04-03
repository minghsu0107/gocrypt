package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"golang.org/x/sync/errgroup"

	"github.com/minghsu0107/gocrypt/pkg/api"
	"github.com/minghsu0107/gocrypt/pkg/config"
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
	// Run: func(cmd *cobra.Command, args []string) {},
	RunE: func(cmd *cobra.Command, args []string) error {
		conf, err := config.NewConfig()
		if err != nil {
			return err
		}
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
			return allcoin.DisplayAllCoins(ctx, conf, dataChannel, &sendData)
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
	// initConfig is called just before executing rootCmd.RunE
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.gocrypt.yaml)")
	rootCmd.PersistentFlags().StringP("currency-init-unit", "i", "", "initial currency unit")

	// local flags
	rootCmd.Flags().StringP("user", "u", "", "root user")

	// bind global flags to viper
	viper.BindPFlag("currency.initUnit", rootCmd.PersistentFlags().Lookup("currency-init-unit"))

	// bind local flags to viper
	viper.BindPFlag("root.user", rootCmd.Flags().Lookup("user"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.SetConfigType("yaml")

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

	// read in environment variables that match
	// eg. GOCRYPT_CURRENCY_INITUNIT=new-taiwan-dollar ./gocrypt
	viper.SetEnvPrefix("GOCRYPT")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
}
