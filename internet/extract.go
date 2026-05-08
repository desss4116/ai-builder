package internet

import (
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ExtractCleanText(html string) string {

	doc, err := goquery.NewDocumentFromReader(
		strings.NewReader(html),
	)

	if err != nil {
		return ""
	}

	doc.Find("script").Remove()
	doc.Find("style").Remove()
	doc.Find("noscript").Remove()
	doc.Find("svg").Remove()
	doc.Find("header").Remove()
	doc.Find("footer").Remove()
	doc.Find("nav").Remove()

	text := doc.Text()

	text = strings.ReplaceAll(text, "\n", " ")
	text = strings.ReplaceAll(text, "\t", " ")

	re := regexp.MustCompile(`\s+`)
	text = re.ReplaceAllString(text, " ")

	if len(text) > 3000 {
		text = text[:3000]
	}

	return strings.TrimSpace(text)
}
