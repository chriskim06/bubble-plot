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
		Container:  lipgloss.NewStyle().Border(lipgloss.NormalBorder()).Margin(1).Padding(0, 1),
		Title:      lipgloss.NewStyle().Bold(true),
		AxisColor:  c,
		LabelColor: c,
	}
}

type Option func(*Model)

func WithContainerStyle(style lipgloss.Style) Option {
	return func(m *Model) {
		m.Styles.Container = style
	}
}

func WithTitleStyle(style lipgloss.Style) Option {
	return func(m *Model) {
		m.Styles.Title = style
	}
}

func WithShowTitle(show bool) Option {
	return func(m *Model) {
		m.showTitle = show
	}
}
