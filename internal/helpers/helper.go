package helpers

import (
	"fmt"
	"strings"
	"time"

	"github.com/clarketm/json"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"github.com/goombaio/namegenerator"
	"github.com/rs/zerolog/log"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// MDToHTML converts markdown to html
func MDToHTML(md string) string {
	// create Markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse([]byte(md))

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}

	renderer := html.NewRenderer(opts)

	b := markdown.Render(doc, renderer)
	return string(b)
}

// RandName creates a random name
func RandName(withHyphen bool) string {
	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)

	name := nameGenerator.Generate()

	if !withHyphen {
		name = strings.ReplaceAll(name, "-", "")
	}
	return name
}

// RandFromName creates a random name from the given name
// this is used to create a random name for the weaviate class
func RandFromName(name string, removeWhiteSpace bool) string {
	name = strings.ReplaceAll(name, "-", "")
	//goland:noinspection SpellCheckingInspection
	caser := cases.Title(language.English)
	name = fmt.Sprintf("%s%s", caser.String(name), RandName(false))

	if removeWhiteSpace {
		name = strings.ReplaceAll(name, " ", "")
	}

	return name
}

// MustMarshalJSON marshals the given value to json
func MustMarshalJSON(v any) []byte {
	b, err := json.Marshal(v)
	if err != nil {
		log.Err(err).Msg("error marshalling json")
		return nil
	}
	return b
}

// MustUnmarshalJSON unmarshals the given json data to the given value
func MustUnmarshalJSON(data []byte, v any) {
	err := json.Unmarshal(data, v)
	if err != nil {
		log.Err(err).Msg("error unmarshalling json")
	}
}
