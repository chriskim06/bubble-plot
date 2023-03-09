package plot

import "github.com/charmbracelet/lipgloss"

type Styles struct {
	Container  lipgloss.Style
	Title      lipgloss.Style
	AxisColor  int
	LabelColor int
	LineColors []int
}

func NewDefaultStyles() Styles {
	c := 231
	if !lipgloss.HasDarkBackground() {
		c = 0
	}
	return Styles{
		Container:  lipgloss.NewStyle().Border(lipgloss.NormalBorder()),
		Title:      lipgloss.NewStyle().Bold(true),
		AxisColor:  c,
		LabelColor: c,
	}
}
