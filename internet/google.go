package internet

import (
	"context"
	"strings"

	"github.com/chromedp/chromedp"
)

func GoogleSearch(query string) string {

	ctx, cancel :=
		chromedp.NewContext(context.Background())

	defer cancel()

	var body string

	url :=
		"https://www.google.com/search?q=" +
			query

	err := chromedp.Run(
		ctx,

		chromedp.Navigate(url),

		chromedp.WaitVisible(`body`),

		chromedp.Text(`body`, &body),
	)

	if err != nil {
		return ""
	}

	body = strings.TrimSpace(body)

	if len(body) > 3000 {
		body = body[:3000]
	}

	return body
}
