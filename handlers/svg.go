package handlers

import (
	"encoding/base64"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ConvertSVGToImg(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatalln("Error parsing HTML:", err)
		panic(err)
	}

	doc.Find("svg").Each(func(_ int, s *goquery.Selection) {
		outerHTML, err := goquery.OuterHtml(s)
		if err != nil {
			log.Fatalln("Error parsing HTML:", err)
			panic(err)
		}
		encodedSVG := base64.StdEncoding.EncodeToString([]byte(outerHTML))
		s.ReplaceWithHtml("<img src=\"data:image/svg+xml;base64," + encodedSVG + "\" />")
	})

	output, err := doc.Html()
	if err != nil {
		log.Fatalln("Error parsing HTML:", err)
		panic(err)
	}

	return output
}
