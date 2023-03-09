package plot

import (
	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	if m.canvas == nil {
		return ""
	}
	sections := []string{}
	if m.showTitle && m.Title != "" {
		title := m.Styles.Title.Render(m.Title)
		sections = append(sections, title)
	}
	sections = append(sections, m.canvas.Plot(m.data))
	return m.Styles.Container.Render(lipgloss.JoinVertical(lipgloss.Left, sections...))
}
