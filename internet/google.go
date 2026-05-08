package internet

import (
	"context"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

func GoogleSearch(query string) string {

	/*
	   CREATE CONTEXT
	*/

	opts := append(
		chromedp.DefaultExecAllocatorOptions[:],

		chromedp.Flag("headless", true),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("no-sandbox", true),

		chromedp.UserAgent(
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0 Safari/537.36",
		),
	)

	allocCtx, cancel :=
		chromedp.NewExecAllocator(
			context.Background(),
			opts...,
		)

	defer cancel()

	ctx, cancel :=
		chromedp.NewContext(allocCtx)

	defer cancel()

	ctx, cancel =
		context.WithTimeout(
			ctx,
			20*time.Second,
		)

	defer cancel()

	/*
	   RESULT
	*/

	var body string

	searchURL :=
		"https://www.google.com/search?q=" + query

	/*
	   RUN
	*/

	err := chromedp.Run(
		ctx,

		chromedp.Navigate(searchURL),

		chromedp.Sleep(3*time.Second),

		chromedp.WaitVisible(`body`),

		chromedp.Text(`body`, &body),
	)

	if err != nil {
		return ""
	}

	/*
	   CLEAN
	*/

	body = strings.TrimSpace(body)

	if len(body) > 4000 {
		body = body[:4000]
	}

	return body
}
