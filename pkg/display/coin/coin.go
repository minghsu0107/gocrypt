package coin

import (
	"context"
	"fmt"
	"strconv"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/minghsu0107/gocrypt/pkg/api"
	"github.com/minghsu0107/gocrypt/pkg/config"
	uw "github.com/minghsu0107/gocrypt/pkg/display/utilitywidgets"
	"github.com/minghsu0107/gocrypt/pkg/utils"
	"github.com/minghsu0107/gocrypt/pkg/widgets"
)

type DisplayType string

var (
	MainDisplay      DisplayType = "main"
	PortfolioDisplay DisplayType = "portfolio"
)

// DisplayCoin displays the per coin values and details along with a favourites table. It uses the same uiEvents channel as the root page
func DisplayCoin(
	ctx context.Context,
	conf *config.Config,
	displayType DisplayType,
	id string,
	coinIDs api.CoinIDMap,
	intervalChannel chan string,
	dataChannel chan api.CoinData,
	priceChannel chan string,
	uiEvents <-chan ui.Event) error {

	defer ui.Clear()

	// Init Coin page
	page := newCoinPage()

	// Currency table
	currencyWidget := uw.NewCurrencyPage()

	currencyID := utils.GetCurrencyID()
	currencyID, currency, currencyVal := currencyWidget.Get(currencyID)

	// variables for graph interval
	changeInterval := "24 Hours"
	changeIntervalWidget := uw.NewChangeIntervalPage()

	// Selection of default table
	selectedTable := page.ExplorerTable
	selectedTable.ShowCursor = true
	utilitySelected := uw.None

	// variables to sort favourites table
	favSortIdx := -1
	favSortAsc := false
	favHeader := []string{
		"Symbol",
		fmt.Sprintf("Price (%s)", currency),
	}

	// Initialise portfolio
	var user string
	switch displayType {
	case MainDisplay:
		user = conf.Root.User
	case PortfolioDisplay:
		user = conf.Portfolio.User
	default:
		user = conf.Root.User
	}
	favourites := utils.GetFavourites(user)
	portfolioMap := utils.GetPortfolio(user)
	defer func() {
		utils.SaveMetadata(favourites, currencyID, user, portfolioMap)
	}()

	// Initiliase Portfolio Table
	portfolioTable := uw.NewPortfolioPage()

	// Initialise help menu
	help := widgets.NewHelpMenu()
	help.SelectHelpMenu("COIN")

	// UpdateUI to refresh UI
	updateUI := func() {
		// Get Terminal Dimensions
		w, h := ui.TerminalDimensions()

		// Adjust Suuply chart Bar graph values
		page.SupplyChart.BarGap = ((w / 3) - (2 * page.SupplyChart.BarWidth)) / 2

		page.Grid.SetRect(0, 0, w, h)

		// Clear UI
		ui.Clear()

		// Render required widgets
		switch utilitySelected {
		case uw.Help:
			help.Resize(w, h)
			ui.Render(help)
		case uw.Portfolio:
			portfolioTable.Resize(w, h)
			ui.Render(portfolioTable)
		case uw.Currency:
			currencyWidget.Resize(w, h)
			ui.Render(currencyWidget)
		case uw.Change:
			changeIntervalWidget.Resize(w, h)
			ui.Render(changeIntervalWidget)
		default:
			ui.Render(page.Grid)
		}
	}

	// Render empty UI
	updateUI()

	// Create ticker to periodically refresh UI
	t := time.NewTicker(time.Duration(1) * time.Second)
	tick := t.C

	previousKey := ""

	for {
		select {
		case <-ctx.Done(): // Context cancelled, exit
			return ctx.Err()

		case e := <-uiEvents: // keyboard events
			switch e.ID {
			case "<Escape>", "q", "<C-c>":
				if utilitySelected != uw.None {
					utilitySelected = uw.None
					selectedTable = page.ExplorerTable
					selectedTable.ShowCursor = true
					updateUI()
				} else {
					return fmt.Errorf("UI Closed")
				}

			case "<Resize>":
				updateUI()

			case "?":
				selectedTable.ShowCursor = false
				selectedTable = help.Table
				selectedTable.ShowCursor = true
				utilitySelected = uw.Help
				updateUI()

			case "c":
				if utilitySelected == uw.None {
					selectedTable.ShowCursor = false
					selectedTable = currencyWidget.Table
					selectedTable.ShowCursor = true
					currencyWidget.UpdateRows(false)
					utilitySelected = uw.Currency
				}

			case "C":
				if utilitySelected == uw.None {
					selectedTable.ShowCursor = false
					selectedTable = currencyWidget.Table
					selectedTable.ShowCursor = true
					currencyWidget.UpdateRows(true)
					utilitySelected = uw.Currency
				}

			case "d":
				if utilitySelected == uw.None {
					selectedTable.ShowCursor = false
					selectedTable = changeIntervalWidget.Table
					selectedTable.ShowCursor = true
					utilitySelected = uw.Change
				}

			case "f":
				if utilitySelected == uw.None {
					selectedTable.ShowCursor = false
					selectedTable = page.FavouritesTable
					selectedTable.ShowCursor = true
				}

			case "F":
				if utilitySelected == uw.None {
					selectedTable.ShowCursor = false
					selectedTable = page.ExplorerTable
					selectedTable.ShowCursor = true
				}

			case "P":
				if utilitySelected == uw.None {
					selectedTable.ShowCursor = false
					selectedTable = portfolioTable.Table
					selectedTable.ShowCursor = true
					portfolioTable.UpdateRows(portfolioMap, currency, currencyVal)
					utilitySelected = uw.Portfolio
				}

			// Navigations
			case "j", "<Down>":
				selectedTable.ScrollDown()

			case "k", "<Up>":
				selectedTable.ScrollUp()

			case "<C-d>":
				selectedTable.ScrollHalfPageDown()

			case "<C-u>":
				selectedTable.ScrollHalfPageUp()

			case "<C-f>":
				selectedTable.ScrollPageDown()

			case "<C-b>":
				selectedTable.ScrollPageUp()

			case "g":
				if previousKey == "g" {
					selectedTable.ScrollTop()
				}

			case "<Home>":
				selectedTable.ScrollTop()

			case "G", "<End>":
				selectedTable.ScrollBottom()

			// Actions
			case "<Enter>":
				switch utilitySelected {
				case uw.Change:
					// Update Graph Durations
					if changeIntervalWidget.SelectedRow < len(changeIntervalWidget.Rows) {
						row := changeIntervalWidget.Rows[changeIntervalWidget.SelectedRow]

						// Get newer selected duration
						changeInterval = row[0]
						newChangeInterval := uw.IntervalMap[changeInterval]

						// Empty current graph
						page.ValueGraph.Data["Value"] = []float64{}

						// Send Updated Interval
						intervalChannel <- newChangeInterval
					}
					utilitySelected = uw.None

				case uw.Currency:

					// Update Currency
					if currencyWidget.SelectedRow < len(currencyWidget.Rows) {
						row := currencyWidget.Rows[currencyWidget.SelectedRow]

						// Get currency and rate
						currencyID = row[0]
						currencyID, currency, currencyVal = currencyWidget.Get(currencyID)

						// Update currency fields
						favHeader[1] = fmt.Sprintf("Price (%s)", currency)
					}
					utilitySelected = uw.None
				}

				if utilitySelected == uw.None {
					selectedTable.ShowCursor = false
					selectedTable = page.ExplorerTable
					selectedTable.ShowCursor = true
				}

			case "e":
				switch utilitySelected {
				case uw.Portfolio:
					id := ""
					symbol := ""

					// Get symbol
					if portfolioTable.SelectedRow < len(portfolioTable.Rows) {
						row := portfolioTable.Rows[portfolioTable.SelectedRow]
						symbol = row[1]
					}

					// Get ID from symbol
					id = coinIDs[symbol].CoinGeckoID

					if id != "" {
						// Draw Edit Box and get new amount
						inputStr := widgets.DrawEdit(uiEvents, symbol)
						amt, err := strconv.ParseFloat(inputStr, 64)

						// Update amount
						if err == nil {
							if amt > 0 {
								portfolioMap[id] = amt
							} else {
								delete(portfolioMap, id)
							}
						}
					}

					portfolioTable.UpdateRows(portfolioMap, currency, currencyVal)
				}
			}

			if utilitySelected == uw.None {
				switch selectedTable {
				case page.FavouritesTable:
					switch e.ID {
					// Sort Ascending
					case "1", "2":
						idx, _ := strconv.Atoi(e.ID)
						favSortIdx = idx - 1
						page.FavouritesTable.Header = append([]string{}, favHeader...)
						page.FavouritesTable.Header[favSortIdx] = favHeader[favSortIdx] + " " + utils.UpArrow
						favSortAsc = true
						utils.SortData(page.FavouritesTable.Rows, favSortIdx, favSortAsc, "FAVOURITES")

					// Sort Descending
					case "<F1>", "<F2>":
						page.FavouritesTable.Header = append([]string{}, favHeader...)
						idx, _ := strconv.Atoi(e.ID[2:3])
						favSortIdx = idx - 1
						page.FavouritesTable.Header[favSortIdx] = favHeader[favSortIdx] + " " + utils.DownArrow
						favSortAsc = false
						utils.SortData(page.FavouritesTable.Rows, favSortIdx, favSortAsc, "FAVOURITES")
					}
				}
			}

			updateUI()
			if previousKey == "g" {
				previousKey = ""
			} else {
				previousKey = e.ID
			}

		case data := <-priceChannel:
			// Update live price
			if data == "NA" {
				if utilitySelected == uw.None {
					page.PriceBox.Rows[0][0] = data
				}
			} else {
				p, _ := strconv.ParseFloat(data, 64)
				if utilitySelected == uw.None {
					page.PriceBox.Rows[0][0] = fmt.Sprintf("%.2f", p/currencyVal)
					ui.Render(page.PriceBox)
				}
			}

		case data := <-dataChannel:
			switch data.Type {

			case "FAVOURITES":
				// Update favorites table
				rows := [][]string{}
				for symbol, price := range data.Favourites {
					p := fmt.Sprintf("%.2f", price/currencyVal)
					rows = append(rows, []string{symbol, p})
				}
				page.FavouritesTable.Header[1] = fmt.Sprintf("Price (%s)", currency)
				page.FavouritesTable.Rows = rows

			case "HISTORY":
				// Update History graph
				price := data.PriceHistory

				// Set value, min & max price
				page.ValueGraph.Data["Value"] = price
				value := (price[len(price)-1] + data.MinPrice) / currencyVal

				page.ValueGraph.Labels["Value"] = fmt.Sprintf("%.2f %s", value, currency)
				page.ValueGraph.Labels["Max"] = fmt.Sprintf("%.2f %s", data.MaxPrice/currencyVal, currency)
				page.ValueGraph.Labels["Min"] = fmt.Sprintf("%.2f %s", data.MinPrice/currencyVal, currency)

				// Update Graph title
				page.ValueGraph.Title = fmt.Sprintf(" Value History (%s) ", changeInterval)

			case "DETAILS":
				// Update Details table
				page.DetailsTable.Header = []string{"Name", data.Details.Name}

				marketCapVals, units := utils.RoundValues(data.Details.MarketCap, 0)
				marketCap := fmt.Sprintf("%.2f %s %s", marketCapVals[0]/currencyVal, units, currency)

				ATHVals, units := utils.RoundValues(data.Details.ATH, 0)
				ATH := fmt.Sprintf("%.2f %s %s", ATHVals[0]/currencyVal, units, currency)

				ATLVals, units := utils.RoundValues(data.Details.ATL, 0)
				ATL := fmt.Sprintf("%.2f %s %s", ATLVals[0]/currencyVal, units, currency)

				TotalVolVals, units := utils.RoundValues(data.Details.TotalVolume, 0)
				TotalVolume := fmt.Sprintf("%.2f %s %s", TotalVolVals[0]/currencyVal, units, currency)

				rows := [][]string{
					{"Symbol", data.Details.Symbol},
					{"Rank", data.Details.Rank},
					{"BlockTime (min)", data.Details.BlockTime},
					{"MarketCap", marketCap},
					{"ATH", ATH},
					{"ATHDate", data.Details.ATHDate},
					{"ATL", ATL},
					{"ATLDate", data.Details.ATLDate},
					{"TotalVolume", TotalVolume},
					{"LastUpdate", data.Details.LastUpdate},
				}

				page.DetailsTable.Rows = rows

				// Update 24 High/Low
				page.PriceBox.Rows[0][1] = fmt.Sprintf("%.2f", data.Details.High24/currencyVal)
				page.PriceBox.Rows[0][2] = fmt.Sprintf("%.2f", data.Details.Low24/currencyVal)
				page.PriceBox.Title = fmt.Sprintf(" Live Price (%s) ", currency)

				// Get Change Percents
				page.ChangesTable.Rows = data.Details.ChangePercents

				// Get supply and Max supply
				supply := data.Details.CurrentSupply
				maxSupply := data.Details.TotalSupply

				supplyVals, units := utils.RoundValues(supply, maxSupply)
				page.SupplyChart.Data = supplyVals
				page.SupplyChart.Title = fmt.Sprintf(" Supply (%s) ", units)

				// Get Explorers
				page.ExplorerTable.Rows = data.Details.Explorers

			}

			// Sort favourites table
			if favSortIdx != -1 {
				utils.SortData(page.FavouritesTable.Rows, favSortIdx, favSortAsc, "FAVOURITES")

				if favSortAsc {
					page.FavouritesTable.Header[favSortIdx] = favHeader[favSortIdx] + " " + utils.UpArrow
				} else {
					page.FavouritesTable.Header[favSortIdx] = favHeader[favSortIdx] + " " + utils.DownArrow
				}
			} else {
				utils.SortData(page.FavouritesTable.Rows, 0, true, "FAVOURITES")
			}

		case <-tick: // Refresh UI
			updateUI()
		}
	}
}
