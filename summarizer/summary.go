package summarizer

import (
	"strings"
)

func Summarize(query string, texts []string) string {

	var final []string

	for _, t := range texts {

		if len(t) < 100 {
			continue
		}

		if len(t) > 700 {
			t = t[:700]
		}

		final = append(final, t)
	}

	if len(final) == 0 {
		return "❌ Информация не найдена."
	}

	result := strings.Join(final, "\n\n")

	if len(result) > 1800 {
		result = result[:1800]
	}

	return "🌐 Ответ:\n\n" + result
}
