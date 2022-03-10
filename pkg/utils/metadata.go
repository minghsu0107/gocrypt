package utils

import (
	"encoding/json"
	"os"

	"github.com/spf13/viper"
)

// Metadata holds persistent information to be stored to disk
type Metadata struct {
	Currency      string                `json:"currency"`
	FavouritesMap map[string]Favourites `json:"favourites_map"`
	PortfolioMap  map[string]Portfolio  `json:"portfolio_map"`
}

type Favourites map[string]bool
type Portfolio map[string]float64

// Currency holds currency data when fetched from CoinCap
type Currency struct {
	ID             string `json:"id"`
	Symbol         string `json:"symbol"`
	CurrencySymbol string `json:"currencySymbol"`
	Type           string `json:"type"`
	RateUSD        string `json:"rateUsd"`
}

// AllCurrencyData holds data of a group of currencies fetched from CoinCap
type AllCurrencyData struct {
	Data      []Currency `json:"data"`
	Timestamp uint       `json:"timestamp"`
}

// GetFavourites reads stored favourite coin details from
// ~/.gocrypt-data.json and returns a map.
func GetFavourites(user string) map[string]bool {
	metadata := Metadata{}

	// Get home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return map[string]bool{}
	}

	// Check if metadata file exists
	configPath := homeDir + "/.gocrypt-data.json"
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return map[string]bool{}
	}

	// Open file
	configFile, err := os.Open(configPath)
	if err != nil {
		return map[string]bool{}
	}

	// Read content
	err = json.NewDecoder(configFile).Decode(&metadata)
	if err != nil {
		return map[string]bool{}
	}

	favourites, ok := metadata.FavouritesMap[user]
	if !ok {
		return map[string]bool{}
	}

	if len(favourites) > 0 {
		return favourites
	}

	return map[string]bool{}
}

// GetPortfolio reads stored portfolio details from
// ~/.gocrypt-data.json and returns a map.
func GetPortfolio(portfolioUser string) map[string]float64 {
	metadata := Metadata{}

	// Get home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return map[string]float64{}
	}

	// Check if metadta file exists
	configPath := homeDir + "/.gocrypt-data.json"
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return map[string]float64{}
	}

	// Open file
	configFile, err := os.Open(configPath)
	if err != nil {
		return map[string]float64{}
	}

	// Read content
	err = json.NewDecoder(configFile).Decode(&metadata)
	if err != nil {
		return map[string]float64{}
	}

	portfolio, ok := metadata.PortfolioMap[portfolioUser]
	if !ok {
		return map[string]float64{}
	}
	if len(portfolio) > 0 {
		return portfolio
	}

	return map[string]float64{}
}

// GetCurrencyID returns the currencyID stored from metadata
func GetCurrencyID() string {
	metadata := Metadata{}
	initCurrency := viper.GetViper().GetString("currency.initUnit")
	// Get home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return initCurrency
	}

	// Check if metadta file exists
	configPath := homeDir + "/.gocrypt-data.json"
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return initCurrency
	}

	// Open file
	configFile, err := os.Open(configPath)
	if err != nil {
		return initCurrency
	}

	// Read content
	err = json.NewDecoder(configFile).Decode(&metadata)
	if err != nil {
		return initCurrency
	}

	return metadata.Currency
}

// SaveMetadata exports favourites, currency and portfolio to disk.
// Data is saved on ~/.gocrypt-data.json
func SaveMetadata(favourites map[string]bool, currency string, user string, portfolio map[string]float64) error {
	var metadata Metadata
	// Get Home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	// Check if metadata file exists
	configPath := homeDir + "/.gocrypt-data.json"
	_, err = os.Stat(configPath)
	if os.IsNotExist(err) {
		favouritesMap := make(map[string]Favourites)
		favouritesMap[user] = favourites

		portfolioMap := make(map[string]Portfolio)
		portfolioMap[user] = portfolio

		metadata = Metadata{
			Currency:      currency,
			FavouritesMap: favouritesMap,
			PortfolioMap:  portfolioMap,
		}
	} else {
		// Open file
		configFile, err := os.Open(configPath)
		if err != nil {
			return err
		}

		// Read content
		err = json.NewDecoder(configFile).Decode(&metadata)
		if err != nil {
			return err
		}

		// Set data
		metadata.Currency = currency
		metadata.FavouritesMap[user] = favourites
		metadata.PortfolioMap[user] = portfolio
	}

	return writeMetaData(homeDir, &metadata)
}

func writeMetaData(homeDir string, metadata *Metadata) error {
	// configPath and hidden path are used explicitly because we
	// get a permission denied error on trying to write/create
	// to a hidden file
	configPath := homeDir + "/gocrypt-data.json"
	hiddenPath := homeDir + "/.gocrypt-data.json"

	data, err := json.MarshalIndent(metadata, "", "\t")
	if err != nil {
		return err
	}

	// Write to file
	err = os.WriteFile(configPath, data, 0666)
	if err != nil {
		return err
	}

	// Hide file
	err = os.Rename(configPath, hiddenPath)
	if err != nil {
		return err
	}
	return nil
}
