package main

import (
	"flag"
	"log"

	"github.com/burgr033/flashMe/internal/data"
	ui "github.com/burgr033/flashMe/internal/ui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	stringFlag := flag.String("cards", "cards.txt", "File path to cards.txt")
	cards, err := data.ParseFlashcards(*stringFlag)
	if err != nil {
		log.Fatalf("Error:%v\n", err)
	}

	p := tea.NewProgram(ui.InitialModel(cards))
	if _, err := p.Run(); err != nil {
		log.Fatalf("Error:%v\n", err)
	}
}
