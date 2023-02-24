package plot

import tea "github.com/charmbracelet/bubbletea"

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.SetSize(msg)
	}
	return m, nil
}

func (m *Model) SetSize(msg tea.WindowSizeMsg) {
	//     h, v := m.Styles.MainStyle.GetFrameSize()
	m.Width = msg.Width
	m.Height = msg.Height
}
