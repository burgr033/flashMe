package types

type FlashCard struct {
	Category string
	Question string `csv:"question"`
	Answer   string `csv:"answer"`
	Rating   int
}
