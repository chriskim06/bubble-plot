package plot

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/chriskim06/drawille-go"
)

type GraphUpdateMsg struct {
	Data [][]float64
}

func GraphUpdateCmd(data [][]float64) tea.Cmd {
	return func() tea.Msg {
		return GraphUpdateMsg{Data: data}
	}
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.SetSize(msg)
	case GraphUpdateMsg:
		m.data = msg.Data
	}
	return m, nil
}

func (m *Model) SetSize(msg tea.WindowSizeMsg) {
	m.Width = msg.Width
	m.Height = msg.Height
	m.Styles.Container.MaxWidth(m.Width).Width(m.Width)
	m.Styles.Container.MaxHeight(m.Height).Height(m.Height)
	h, v := m.Styles.Container.GetFrameSize()
	if m.showTitle && m.Title != "" {
		v += m.Styles.Title.GetVerticalFrameSize()
	}
	canvas := drawille.NewCanvas(m.Width-h, m.Height-v)
	canvas.AxisColor = drawille.SeaGreen
	canvas.LineColors = []drawille.Color{drawille.Red, drawille.SeaGreen}
	m.canvas = &canvas
}
