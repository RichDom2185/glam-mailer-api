package handlers

import (
	"log"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	VAR_CALL_REGEXP = `var\((--[^)]+)\)`
)

// Precondition: All CSS styles are inlined
func SubstituteCSSVariables(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatalln("Error parsing HTML:", err)
		panic(err)
	}

	// Extract CSS variable names and values
	doc.Find("[style]").Each(func(_ int, s *goquery.Selection) {
		styleAttribute, _ := s.Attr("style")
		styleProperties := strings.Split(styleAttribute, ";")

		var computedStyles []string
		cssVariableLookup := make(map[string]string)
		for _, prop := range styleProperties {
			prop = strings.TrimSpace(prop)
			if prop == "" {
				continue
			}

			// TODO: Investigate if there is a better way to parse this
			isCSSVariable := strings.HasPrefix(prop, "--") && strings.Contains(prop, ":")
			if !isCSSVariable {
				computedStyles = append(computedStyles, prop)
				continue
			}
			parts := strings.SplitN(prop, ":", 2)
			name := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			cssVariableLookup[name] = value
		}

		// Substitute CSS variables for that tag
		// Assumes no fallback value e.g. var(--color, #000000)
		// TODO: Support fallback values
		regex := regexp.MustCompile(VAR_CALL_REGEXP)
		for i, computedStyle := range computedStyles {
			computedStyle = regex.ReplaceAllStringFunc(computedStyle, func(match string) string {
				// Capture group 1, as defined in VAR_CALL
				propName := regex.FindStringSubmatch(match)[1]
				// Assumes variable lookups always succeed
				// TODO: Handle the case where it fails
				return cssVariableLookup[propName]
			})
			computedStyles[i] = computedStyle
		}

		// Replace the entire style attribute
		// This removes the now unnecesary variable declarations
		s.SetAttr("style", strings.Join(computedStyles, ";"))
	})

	output, err := doc.Html()
	if err != nil {
		log.Fatalln("Error parsing HTML:", err)
		panic(err)
	}

	return output
}
