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

	mainText := texts[0]

	// LANGUAGE

	isRussian :=
		DetectRussian(query)

	if isRussian {

		return "🌐 Ответ:\n\n" + mainText
	}

	return "🌐 Answer:\n\n" + mainText
}
