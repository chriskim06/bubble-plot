package plot

import "github.com/charmbracelet/lipgloss"

type Styles struct {
	Width  int
	Height int

	MainStyle lipgloss.Style
	TitleBar  lipgloss.Style
	Title     lipgloss.Style
}

func NewDefaultStyles() Styles {
	return Styles{}
}
