package crawler

import (
	"context"

	"github.com/chromedp/chromedp"
)

func RenderPage(url string) string {

	ctx, cancel := chromedp.NewContext(context.Background())

	defer cancel()

	var html string

	chromedp.Run(
		ctx,
		chromedp.Navigate(url),
		chromedp.OuterHTML("html", &html),
	)

	return html
}
