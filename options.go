package plot

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/chriskim06/drawille-go"
)

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

func WithAxisColor(color int) Option {
	return func(m *Model) {
		m.Styles.AxisColor = color
	}
}

func WithLabelColor(color int) Option {
	return func(m *Model) {
		m.Styles.LabelColor = color
	}
}

func WithLineColors(colors []int) Option {
	return func(m *Model) {
		m.Styles.LineColors = colors
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

func WithData(data [][]float64) Option {
	return func(m *Model) {
		m.data = data
	}
}

func WithHorizontalLabels(labels []string) Option {
	return func(m *Model) {
		m.horizontalLabels = labels
	}
}

func WithDimensions(width, height int) Option {
	return func(m *Model) {
		c := drawille.NewCanvas(width, height)
		m.Width = width
		m.Height = height
		m.canvas = &c
	}
}
