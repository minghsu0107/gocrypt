package utilitywidgets

import (
	ui "github.com/gizak/termui/v3"
	"github.com/minghsu0107/gocrypt/pkg/widgets"
)

var intervalRows = [][]string{{"24 Hours"}, {"7 Days"}, {"14 Days"}, {"30 Days"}, {"90 Days"}, {"180 Days"}, {"1 Year"}, {"5 Years"}}

// IntervalMap maps given interval string to format required by CoinGecko API
var IntervalMap = map[string]string{
	"24 Hours": "24hr",
	"7 Days":   "7d",
	"14 Days":  "14d",
	"30 Days":  "30d",
	"90 Days":  "90d",
	"180 Days": "180d",
	"1 Year":   "1yr",
	"5 Years":  "5yr",
}

// ChangeIntervalDurationTable holds a table to help user change duration intervals
type ChangeIntervalDurationTable struct {
	*widgets.Table
}

// NewChangeIntervalPage returns a pointer to an instance of
// ChangeIntervalDurationTable
func NewChangeIntervalPage() *ChangeIntervalDurationTable {
	c := &ChangeIntervalDurationTable{
		Table: widgets.NewTable(),
	}

	c.Table.Title = " Select Duration for Coin History Interval"
	c.Table.Header = []string{"Duration"}
	c.Table.Rows = intervalRows
	c.Table.CursorColor = ui.ColorCyan
	c.Table.ShowCursor = true
	c.Table.ColWidths = []int{5}
	c.Table.ColResizer = func() {
		x := c.Table.Inner.Dx()
		c.Table.ColWidths = []int{
			x,
		}
	}
	return c
}

// Resize helps resize the ChangeIntervalDurationTable according to terminal dimensions
func (c *ChangeIntervalDurationTable) Resize(termWidth, termHeight int) {
	textWidth := 50

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
func (c *ChangeIntervalDurationTable) Draw(buf *ui.Buffer) {
	c.Table.Draw(buf)
}
