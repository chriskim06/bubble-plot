package plot

func (m Model) View() string {
	return m.canvas.String()
}
