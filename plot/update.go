package plot

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/chriskim06/drawille-go"
)

type GraphUpdateMsg struct {
	Data [][]float64
}

var d = [][]float64{{150, 150}, {30, 30}}

func GraphUpdateCmd(data [][]float64) tea.Cmd {
	return func() tea.Msg {
		return GraphUpdateMsg{Data: data}
	}
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Title = "test"
		m.SetSize(msg)
	case tea.KeyMsg:
		if msg.String() == "q" {
			return m, tea.Quit
		}
	case GraphUpdateMsg:
		m.data = msg.Data
		return m, m.tickCmd()
	}
	return m, nil
}

func (m Model) tickCmd() tea.Cmd {
	return tea.Tick(3*time.Second, func(t time.Time) tea.Msg {
		d[0] = append(d[0], 150)
		d[1] = append(d[1], float64(t.Second()))
		return GraphUpdateMsg{
			Data: d,
		}
	})
}

func (m *Model) SetSize(msg tea.WindowSizeMsg) {
	m.Width = msg.Width / 2
	m.Height = msg.Height / 2
	m.Styles.Container.MaxWidth(m.Width)
	m.Styles.Container.MaxHeight(m.Height)
	h, v := m.Styles.Container.GetFrameSize()
	if m.showTitle && m.Title != "" {
		v += m.Styles.Title.GetVerticalFrameSize()
	}
	canvas := drawille.NewCanvas(m.Width-h, m.Height-v)
	canvas.LineColors = []drawille.AnsiColor{drawille.Red, drawille.SeaGreen}
	m.canvas = &canvas
}
