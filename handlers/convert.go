package handlers

import (
	"context"
	"log"
	"os"

	"github.com/RichDom2185/go-tailwind"
	"github.com/aymerick/douceur/inliner"
	"github.com/livebud/js"
	v8 "github.com/livebud/js/v8"
)

func GenerateRequiredTailwindStyles(html string) string {
	vm, err := v8.Load(&js.Console{
		Log:   os.Stdout,
		Error: os.Stderr,
	})
	defer vm.Close()
	if err != nil {
		log.Fatalln(err)
	}

	// Generate the CSS for this index.html file
	processor := tailwind.New(vm, false)
	ctx := context.Background()
	css, err := processor.Process(ctx, "_.html", html)
	if err != nil {
		log.Fatalln(err)
	}
	return css
}

func InlineCSS(html string) string {
	inlineStyledHTML, err := inliner.Inline(string(html))
	if err != nil {
		panic("Please fill a bug :)")
	}
	log.Default().Println("Successfully inlined CSS!")
	return inlineStyledHTML
}
