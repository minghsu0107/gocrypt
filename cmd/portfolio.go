package cmd

import (
	"context"

	"github.com/minghsu0107/gocrypt/pkg/api"
	"github.com/minghsu0107/gocrypt/pkg/display/portfolio"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

// portfolioCmd represents the portfolio command
var portfolioCmd = &cobra.Command{
	Use:   "portfolio",
	Short: "Track your portfolio",
	Long:  `The portfolio command helps track your own portfolio in real time`,
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

		// Display UI for portfolio
		eg.Go(func() error {
			return portfolio.DisplayPortfolio(ctx, dataChannel, &sendData)
		})

		if err := eg.Wait(); err != nil {
			if err.Error() != "UI Closed" {
				return err
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(portfolioCmd)
}
