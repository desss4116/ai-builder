package brain

import (
	"ai-builder/rag"
	"strings"
)

func Synthesize(
	query string,
	texts []string,
) string {

	intent :=
		DetectIntent(query)

	// RAG MEMORY

	memoryResults :=
		rag.Retrieve(query)

	texts = append(
		texts,
		memoryResults...,
	)

	// SEMANTIC RANKING

	texts = RankTexts(
		query,
		texts,
	)

	// RECOMMENDATIONS

	if intent == "recommendation" {

		rec := Recommend(query)

		if rec != "" {
			return rec
		}
	}

	if len(texts) == 0 {
		return "❌ Информация не найдена."
	}

	// CLEANUP

	cleaned := []string{}

	for _, t := range texts {

		t = strings.ReplaceAll(
			t,
			"\n",
			" ",
		)

		t = strings.TrimSpace(t)

		if len(t) < 50 {
			continue
		}

		// REMOVE GARBAGE

		bad := []string{
			"youtube",
			"rutube",
			"playlist",
			"sign up",
			"privacy",
			"cookie",
			"blocked",
			"jump to content",
		}

		skip := false

		lower :=
			strings.ToLower(t)

		for _, b := range bad {

			if strings.Contains(
				lower,
				b,
			) {
				skip = true
			}
		}

		if skip {
			continue
		}

		cleaned = append(
			cleaned,
			t,
		)
	}

	if len(cleaned) == 0 {
		return "❌ Нормальная информация не найдена."
	}

	// BUILD AI RESPONSE

	final := "🌐 Ответ:\n\n"

	used := map[string]bool{}

	for _, t := range cleaned {

		parts :=
			strings.Split(
				t,
				".",
			)

		for _, p := range parts {

			p = strings.TrimSpace(p)

			if len(p) < 40 {
				continue
			}

			if used[p] {
				continue
			}

			used[p] = true

			final += p + ".\n\n"

			if len(final) > 1200 {
				return final
			}
		}
	}

	return final
}
