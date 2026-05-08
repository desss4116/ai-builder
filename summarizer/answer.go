package summarizer

import (
	"strings"
)

func BuildAnswer(text string) string {

	text = strings.TrimSpace(text)

	if len(text) > 1200 {
		text = text[:1200]
	}

	return "🌐 Ответ:\n\n" + text
}
