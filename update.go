package plot

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/chriskim06/drawille-go"
)

type GraphUpdateMsg struct {
	Data   [][]float64
	Labels []string
}

func GraphUpdateCmd(data [][]float64, labels []string) tea.Cmd {
	return func() tea.Msg {
		return GraphUpdateMsg{
			Data: data,
			Labels: labels,
		}
	}
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.SetSize(msg)
	case GraphUpdateMsg:
		m.data = msg.Data
		m.canvas.HorizontalLabels = msg.Labels
	}
	return m, nil
}

func (m *Model) SetSize(msg tea.WindowSizeMsg) {
	m.Width = msg.Width
	m.Height = msg.Height
	m.Styles.Container.Width(m.Width).Height(m.Height)
	canvas := drawille.NewCanvas(m.Width, m.Height)
	if m.MaxDataPoints > 0 {
		canvas.NumDataPoints = m.MaxDataPoints
	}
	canvas.AxisColor = drawille.Color(m.Styles.AxisColor)
	canvas.LabelColor = drawille.Color(m.Styles.LabelColor)
	lineColors := []drawille.Color{}
	for _, c := range m.Styles.LineColors {
		lineColors = append(lineColors, drawille.Color(c))
	}
	canvas.LineColors = lineColors
	if len(m.horizontalLabels) > 0 {
		canvas.HorizontalLabels = m.horizontalLabels
	}
	m.canvas = &canvas
}
