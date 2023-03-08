package main

import (
	"fmt"
	"math"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	plot "github.com/chriskim06/bubble-plot"
)

var x = -1

type model struct {
	data [][]float64
	p    *plot.Model
}

func newmodel() *model {
	pl := plot.New()
	pl.MaxDataPoints = 40
	return &model{
		data: [][]float64{{}, {}},
		p:    pl,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.p.Update(tea.WindowSizeMsg{
			Width:  120,
			Height: 40,
		})
		return m, m.tick()
	case tea.KeyMsg:
		if msg.String() == "q" {
			return m, tea.Quit
		}
	case tickmsg:
		// update data
		if len(m.data[0]) == 40 {
			m.data[0] = m.data[0][1:]
			m.data[1] = m.data[1][1:]
		}
		m.data[0] = append(m.data[0], 3)
		m.data[1] = append(m.data[1], sindata())
		m.p.Update(plot.GraphUpdateMsg{
			Data: m.data,
		})
		return m, m.tick()
	}
	return m, nil
}

func (m model) View() string {
	return m.p.View()
}

type tickmsg struct{}

func (m model) tick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickmsg{}
	})
}

func sindata() float64 {
	x++
	return 2 * math.Sin((math.Pi/9)*float64(x))
}

func main() {
	m := newmodel()
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
