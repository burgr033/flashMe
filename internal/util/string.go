package util

import (
	"log"
	"regexp"

	"github.com/burgr033/latex2unicode/pkg/converter"
)

func ConvertMathString(input string) string {
	re := regexp.MustCompile(`\$(.*?)\$`)
	if !re.MatchString(input) {
		return input
	}

	lc, err := converter.New()
	if err != nil {
		log.Printf("Error loading latex converter: %v\n", err)
		return input
	}

	output := re.ReplaceAllStringFunc(input, func(match string) string {
		inner := match[1 : len(match)-1]

		converted := lc.ConvertMarkdown(inner)

		return "$" + converted + "$"
	})

	return output
}

