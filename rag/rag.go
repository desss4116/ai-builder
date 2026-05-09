package rag

import (
	"ai-builder/memory"
	"ai-builder/vector"
)

func Retrieve(query string) []string {

	var results []string

	memories := memory.Get()

	for _, m := range memories {

		score :=
			vector.Similarity(
				query,
				m,
			)

		if score > 0 {

			results = append(
				results,
				m,
			)
		}
	}

	return results
}
