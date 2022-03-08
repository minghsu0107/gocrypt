package utilitywidgets

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	ui "github.com/gizak/termui/v3"
	"github.com/minghsu0107/gocrypt/pkg/utils"
	"github.com/minghsu0107/gocrypt/pkg/widgets"
)

// SingleCurrency holds data of a currency. Used for API fetches
type SingleCurrency struct {
	ID             string `json:"id"`
	Symbol         string `json:"symbol"`
	CurrencySymbol string `json:"currencySymbol"`
	Type           string `json:"type"`
	RateUSD        string `json:"rateUSD"`
}

// AllCurrencyData holds details of currencies when all are fetched from the API
type AllCurrencyData struct {
	Data      []SingleCurrency `json:"data"`
	Timestamp uint             `json:"timestamp"`
}

// CurrencyTable is a widget used to display currencyies, symbols and rates
type CurrencyTable struct {
	*widgets.Table
	IDMap *CurrencyIDMap
}

// Currency holds information of a single currency, it used to populate currencyIDMaps
type CurrencyValue struct {
	Symbol  string
	RateUSD float64
	Type    string
}

// CurrencyIDMap maps a currency Id to it's symbol and price in USD
type CurrencyIDMap map[string]CurrencyValue

// NewCurrencyIDMap creates and returns an instance of CurrencyIDMap
func NewCurrencyIDMap() CurrencyIDMap {
	c := make(CurrencyIDMap)
	return c
}

// Populate fetches currency rates and populates the map
func (c CurrencyIDMap) Populate() {
	url := "https://api.coincap.io/v2/rates"
	method := "GET"

	client := &http.Client{}

	// Create Request
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return
	}

	// Send Request and get response
	res, err := client.Do(req)
	if err != nil {
		res.Body.Close()
		return
	}

	data := utils.AllCurrencyData{}

	// Read response
	err = json.NewDecoder(res.Body).Decode(&data)
	res.Body.Close()
	if err != nil {
		return
	}

	// Iterate over currencies
	for _, curr := range data.Data {
		currencyID := curr.ID
		rate, err := strconv.ParseFloat(curr.RateUSD, 64)
		if err == nil {

			c[currencyID] = CurrencyValue{
				Symbol:  fmt.Sprintf("%s %s", curr.Symbol, curr.CurrencySymbol),
				RateUSD: rate,
				Type:    curr.Type,
			}
		}
	}
}

// Get returns the symbol and USD rate for a given currency ID
// If the given currency ID does not exist in the Map, values
// for US Dollar are returned
func (c *CurrencyTable) Get(currencyID string) (string, string, float64) {
	if val, ok := (*c.IDMap)[currencyID]; ok {
		return currencyID, val.Symbol, val.RateUSD
	}

	return "united-states-dollar", "USD $", 1

}

// NewCurrencyPage creates, initialises and returns a pointer to an instance of CurrencyTable
func NewCurrencyPage() *CurrencyTable {
	idMap := NewCurrencyIDMap()
	idMap.Populate()

	c := &CurrencyTable{
		Table: widgets.NewTable(),
		IDMap: &idMap,
	}

	c.Table.Title = " Select Currency "
	c.Table.Header = []string{"Currency", "Symbol", "Type", "USD rate"}
	c.Table.CursorColor = ui.ColorCyan
	c.Table.ShowCursor = true
	c.Table.ColWidths = []int{5, 5, 5, 5}
	c.Table.ColResizer = func() {
		x := c.Table.Inner.Dx()
		c.Table.ColWidths = []int{
			4 * x / 10,
			2 * x / 10,
			2 * x / 10,
			2 * x / 10,
		}
	}

	return c
}

// Resize resizes the Currency Table as per given terminal dimensions
func (c *CurrencyTable) Resize(termWidth, termHeight int) {
	textWidth := 80

	textHeight := len(c.Table.Rows) + 3
	x := (termWidth - textWidth) / 2
	y := (termHeight - textHeight) / 2
	if x < 0 {
		x = 0
		textWidth = termWidth
	}
	if y < 0 {
		y = 0
		textHeight = termHeight
	}

	c.Table.SetRect(x, y, textWidth+x, textHeight+y)
}

// Draw puts the required text into the widget
func (c *CurrencyTable) Draw(buf *ui.Buffer) {
	if len(c.Table.Rows) == 0 {
		c.Table.Title = " Unable to fetch currencies, please close and retry "
	} else {
		c.Table.Title = " Select Currency "
	}
	c.Table.Draw(buf)
}

// UpdateRows fetches rates of all currencies and updates them as rows in the table
func (c *CurrencyTable) UpdateRows(allCurrencies bool) {
	currencies := map[string]bool{
		"united-states-dollar":   true,
		"euro":                   true,
		"japanese-yen":           true,
		"british-pound-sterling": true,
		"indian-rupee":           true,
		"australian-dollar":      true,
		"canadian-dollar":        true,
		"chinese-yuan-renminbi":  true,
		"new-taiwan-dollar":      true,
	}

	c.IDMap.Populate()

	rows := make([][]string, 0)

	if allCurrencies {
		// Iterate over all currencies
		for currencyID, currency := range *c.IDMap {
			// Aggregate data
			row := []string{
				currencyID,
				currency.Symbol,
				currency.Type,
				fmt.Sprintf("%.4f", currency.RateUSD),
			}

			rows = append(rows, row)
		}
	} else {
		// Iterate over selected currencies
		for currencyID := range currencies {
			currency := (*c.IDMap)[currencyID]
			// Aggregate data
			row := []string{
				currencyID,
				currency.Symbol,
				currency.Type,
				fmt.Sprintf("%.4f", currency.RateUSD),
			}

			rows = append(rows, row)
		}
	}

	// Update table rows and sort alphabetically
	c.Table.Rows = rows
	utils.SortData(c.Table.Rows, 0, true, "CURRENCY")
}
