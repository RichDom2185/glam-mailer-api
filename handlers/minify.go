package handlers

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func RemoveStyleTags(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatalln("Error parsing HTML:", err)
		panic(err)
	}

	doc.Find("style").Remove()

	output, err := doc.Html()
	if err != nil {
		log.Fatalln("Error parsing HTML:", err)
		panic(err)
	}

	return output
}
