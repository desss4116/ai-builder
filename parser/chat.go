package parser

import (
	"strings"
)

func CleanResponse(text string) string {

	text = strings.TrimSpace(text)

	if len(text) > 1800 {
		text = text[:1800]
	}

	return text
}
