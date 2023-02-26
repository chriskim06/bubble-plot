package plot

import (
	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	if m.canvas == nil {
		return ""
	}
	sections := []string{}
	availableHeight := m.Height
	if m.showTitle && m.Title != "" {
		title := m.Styles.Title.Render(m.Title)
		availableHeight -= lipgloss.Height(title)
		sections = append(sections, title)
	}
	sections = append(sections, m.canvas.Plot(m.data))
	g := lipgloss.JoinVertical(lipgloss.Left, sections...)
	return m.Styles.Container.Render(g)
}
