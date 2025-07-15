package scraper

import (
	"slices"
	"strings"

	"golang.org/x/net/html"
)

var lineBreakTags []string = []string{"div", "br", "li", "ul", "ol"}

func HTMLtoCleanText(htmlStr string) string {
	tokenizer := html.NewTokenizer(strings.NewReader(htmlStr))
	output := strings.Builder{}

	for {
		tt := tokenizer.Next()

		switch tt {
		case html.ErrorToken:
			return output.String()
		case html.TextToken:
			output.WriteString(" " + strings.TrimSpace(string(tokenizer.Text())))
		case html.StartTagToken, html.EndTagToken:
			tagNameBytes, _ := tokenizer.TagName()
			tagName := string(tagNameBytes)
			if slices.Contains(lineBreakTags, tagName) {
				output.WriteString("\n")
			}
		}
	}

}
