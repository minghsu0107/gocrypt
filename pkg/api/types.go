package api

import geckoTypes "github.com/superoo7/go-gecko/v3/types"

// CoinData Holds data pertaining to a single coin.
// This is used to serve per coin details.
// It additionally holds a map of favourite coins.
type CoinData struct {
	Type         string
	PriceHistory []float64
	MinPrice     float64
	MaxPrice     float64
	Details      CoinDetails
	Favourites   map[string]float64
}

// CoinDetails holds information about a coin
type CoinDetails struct {
	Name           string
	Symbol         string
	Rank           string
	BlockTime      string
	MarketCap      float64
	Website        string
	Explorers      [][]string
	ATH            float64
	ATHDate        string
	ATL            float64
	ATLDate        string
	High24         float64
	Low24          float64
	TotalVolume    float64
	ChangePercents [][]string
	TotalSupply    float64
	CurrentSupply  float64
	LastUpdate     string
}

// AssetData is used to hold details of multiple coins and the price history
// of top ranked coins along with their names
type AssetData struct {
	TopCoinData [][]float64
	MaxPrices   []float64
	MinPrices   []float64
	TopCoins    []string
	AllCoinData geckoTypes.CoinsMarket
}

// CoinCapAsset is used to marshal asset data from coinCap APIs
type CoinCapAsset struct {
	ID                string `json:"id"`
	Rank              string `json:"rank"`
	Symbol            string `json:"symbol"`
	Name              string `json:"name"`
	Supply            string `json:"supply"`
	MaxSupply         string `json:"maxSupply"`
	MarketCapUsd      string `json:"marketCapUsd"`
	VolumeUsd24Hr     string `json:"volumeUsd24Hr"`
	PriceUsd          string `json:"priceUsd"`
	ChangePercent24Hr string `json:"changePercent24Hr"`
	Vwap24Hr          string `json:"vwap24Hr"`
	Explorer          string `json:"explorer"`
}

// CoinCapData is used to marshall multiple assets from CoinCap APIs
type CoinCapData struct {
	Data      []CoinCapAsset `json:"data"`
	Timestamp uint           `json:"timestamp"`
}

// CoinID holds the ID of a coin as stored in CoinGecko and CoinCap
type CoinID struct {
	CoinGeckoID string
	CoinCapID   string
}

// CoinIDMap maps a symbol to it's respective ID
type CoinIDMap map[string]CoinID
