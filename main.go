package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

type flashCard struct {
	category string
	question string
	answer   string
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
