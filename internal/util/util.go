package util

import (
	"math/rand"
	"time"

	t "github.com/burgr033/flashMe/internal/types"
)

const minWidth = 64

func ShuffleFlashCards(cards []t.FlashCard) []t.FlashCard {
	randomSource := rand.NewSource(time.Now().UnixNano())
	randomNumberGenerator := rand.New(randomSource)
	for i := len(cards) - 1; i > 0; i-- {
		j := randomNumberGenerator.Intn(i + 1)
		cards[i], cards[j] = cards[j], cards[i]
	}
	return cards
}

func CalulateMaxLines(cards []t.FlashCard) int {
	var highestLineCount int
	for _, card := range cards {
		answerLength := len(card.Answer)
		questionLength := len(card.Question)
		if answerLength > highestLineCount || questionLength > highestLineCount {
			highestLineCount = max(answerLength, questionLength)
		}
	}
	maxLineCount := highestLineCount / minWidth
	return maxLineCount + 1
}
