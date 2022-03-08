package utilitywidgets

import (
	ui "github.com/gizak/termui/v3"
	"github.com/minghsu0107/gocrypt/pkg/widgets"
)

var durationRows = [][]string{{"1 Hour"}, {"24 Hours"}, {"7 Days"}, {"14 Days"}, {"30 Days"}, {"200 Days"}, {"1 Year"}}

// DurationMap maps duration strings to the format required by coinGecko API
var DurationMap = map[string]string{
	"1 Hour":   "1h",
	"24 Hours": "24h",
	"7 Days":   "7d",
	"14 Days":  "14d",
	"30 Days":  "30d",
	"200 Days": "200d",
	"1 Year":   "1y",
}

// ChangePercentageDurationTable holds a table which helps change percentage duration values
type ChangePercentageDurationTable struct {
	*widgets.Table
}

// NewChangePercentPage creates, initialises and returns a pointer to
// an instance of CurrencyTable
func NewChangePercentPage() *ChangePercentageDurationTable {
	c := &ChangePercentageDurationTable{
		Table: widgets.NewTable(),
	}

	c.Table.Title = " Select Duration for Percentage Change "
	c.Table.Header = []string{"Duration"}
	c.Table.Rows = durationRows
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

// Resize helps rresize the ChangePercentageDurationTable according to terminal dimensions
func (c *ChangePercentageDurationTable) Resize(termWidth, termHeight int) {
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
func (c *ChangePercentageDurationTable) Draw(buf *ui.Buffer) {
	c.Table.Draw(buf)
}
