package handlers

import (
	"log"

	"github.com/aymerick/douceur/inliner"
)

func InlineCSS(html string) string {
	inlineStyledHTML, err := inliner.Inline(string(html))
	if err != nil {
		panic("Please fill a bug :)")
	}
	log.Default().Println("Successfully inlined CSS!")
	return inlineStyledHTML
}
