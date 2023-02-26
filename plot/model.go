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

	horizontalScale float64
	canvas          *drawille.Canvas
	showTitle       bool
	data            [][]float64

	horizontalLabelStart float64
	horizontalLabelEnd   float64
}

func New(options ...Option) *Model {
	m := &Model{
		Styles:    NewDefaultStyles(),
		showTitle: true,
	}
	for _, option := range options {
		option(m)
	}
	return m
}

func (m Model) Init() tea.Cmd {
	return m.tickCmd()
}
