package data

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	t "github.com/burgr033/flashMe/internal/types"
)

// ParseFlashcards reads a tab-separated flashcard file and returns a slice of flashcards
func ParseFlashcards(filename string) ([]t.FlashCard, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var cards []t.FlashCard

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Skip empty lines
		if strings.TrimSpace(line) == "" {
			continue
		}

		// Skip comment lines (instructions, etc.)
		if strings.HasPrefix(line, "#") {
			continue
		}

		// Parse tab-separated question/answer
		parts := strings.Split(line, "\t")
		if len(parts) >= 2 {
			card := t.FlashCard{
				Question: strings.TrimSpace(parts[0]),
				Answer:   strings.TrimSpace(parts[1]),
			}
			cards = append(cards, card)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return cards, nil
}
