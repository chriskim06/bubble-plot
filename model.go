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

	horizontalScale  float64
	canvas           *drawille.Canvas
	showTitle        bool
	data             [][]float64
	horizontalLabels []string
}

func New(options ...Option) *Model {
	c := drawille.NewCanvas(0, 0)
	m := &Model{
		Styles:           NewDefaultStyles(),
		showTitle:        true,
		horizontalLabels: []string{},
		canvas:           &c,
	}
	for _, option := range options {
		option(m)
	}
	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}
