package plot

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/chriskim06/drawille-go"
)

type Model struct {
	Width  int
	Height int
	Title  string
	Styles Styles

	canvas drawille.Canvas
}

func New() *Model {
	return &Model{
		canvas: drawille.NewCanvas(),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

// SetData is what sets the graph
func (m *Model) SetData(data [][]float64) {
	maxVal, _ := GetMaxFloat64From2dSlice(data)
	for i, line := range data {
		previousHeight := (line[1] / maxVal) * float64(m.Height-1)
		for j, val := range line[1:] {
			height := (val / maxVal) * float64(m.Height-1) // percentage of the draw area height this point should be displayed
			m.canvas.DrawLine(
				i,
				float64(j)*horizontalScale*2, // need horizontal scale to fit in widget
				(maxVal-previousHeight-1)*4,
				float64(j+1)*horizontalScale*2,
				(maxVal-height-1)*4,
			)
			previousHeight = height
		}
	}
}
