package plot

import "github.com/charmbracelet/lipgloss"

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

func WithMaxDataPoints(amount int) Option {
	return func(m *Model) {
		m.MaxDataPoints = amount
	}
}
