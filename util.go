package main

import (
	"math/rand"
	"time"
)

const minWidth = 64

func shuffleFlashCards(cards []flashCard) []flashCard {
	randomSource := rand.NewSource(time.Now().UnixNano())
	randomNumberGenerator := rand.New(randomSource)
	for i := len(cards) - 1; i > 0; i-- {
		j := randomNumberGenerator.Intn(i + 1)
		cards[i], cards[j] = cards[j], cards[i]
	}
	return cards
}

func calulateMaxLines(cards []flashCard) int {
	var highestLineCount int
	for _, card := range cards {
		answerLength := len(card.answer)
		questionLength := len(card.question)
		if answerLength > highestLineCount || questionLength > highestLineCount {
			if answerLength > questionLength {
				highestLineCount = answerLength
			} else {
				highestLineCount = questionLength
			}
		}
	}
	maxLineCount := highestLineCount / minWidth
	return maxLineCount + 1
}
