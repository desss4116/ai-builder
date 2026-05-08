package brain

import (
	"strings"
)

func Reason(
	query string,
	search string,
	scraped string,
) string {

	text :=
		search + "\n" + scraped

	text =
		SemanticCleanup(text)

	if strings.TrimSpace(text) == "" {

		return "❌ Информация не найдена."
	}

	return "🌐 Ответ:\n\n" + text
}
