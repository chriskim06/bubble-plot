package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/chriskim06/bubble-plot/plot"
)

var x = -1

type model struct {
	data [][]float64
	p plot.Model
}

func newmodel() model {
	pl := plot.New()
	pl.MaxDataPoints = 50
	return model{
		data: [][]float64{{}, {}},
		p: plot.New(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m *model) Update(msg tea.Msg) (model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.p.Update(tea.WindowSizeMsg{
			Width: 120,
			Height: 80,
		})
	case tickmsg:
		// update data
		if len(m.data[0]) == 50 {
			m.data[0] = m.data[0][1:]
			m.data[1] = m.data[1][1:]
		}
		m.data[0] = append(m.data[0], 40)
		m.data[1] = append(m.data[1], sindata())
		return *m, m.tick()
	}
	return *m, nil
}

func (m model) View() string {
	return m.p.Plot(m.data)
}

type tickmsg struct{}

func (m mode) tick() tea.Cmd {
	return tea.Tick(2*time.Second, func(t time.Time) tea.Msg {
		return tickmsg{}
	})
}

func sindata() float64 {
	x++
	return math.Sin((math.Pi/180) * float64(10*x)) + 10
}

func main() {
	if _, err := tea.NewProgram(newmodel()).Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
