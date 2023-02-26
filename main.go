package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/chriskim06/bubble-plot/plot"
)

func main() {
	if _, err := tea.NewProgram(plot.New(), tea.WithAltScreen()).Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
