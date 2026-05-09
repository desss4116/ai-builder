package brain

import (
	"strings"
)

func Synthesize(query string, texts []string) string {

	var final []string

	seen := map[string]bool{}

	for _, t := range texts {

		t = strings.TrimSpace(t)

		if len(t) < 120 {
			continue
		}

		if seen[t] {
			continue
		}

		seen[t] = true

		if len(t) > 500 {
			t = t[:500]
		}

		final = append(final, t)

		if len(final) >= 5 {
			break
		}
	}

	if len(final) == 0 {
		return "❌ Информация не найдена."
	}

	result := strings.Join(
		final,
		"\n\n",
	)

	if len(result) > 1800 {
		result = result[:1800]
	}

	return "🌐 Ответ:\n\n" + result
}
