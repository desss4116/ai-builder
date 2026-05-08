package search

import "strings"

func SemanticSearch(query string) string {

	q := strings.ToLower(query)

	if strings.Contains(q, "фильм") {

		return `
Semantic matches found:

- Interstellar
- Oppenheimer
- Dune Part Two
`
	}

	if strings.Contains(q, "ai") {

		return `
Semantic AI knowledge found:
- machine learning
- neural networks
- reasoning systems
`
	}

	return `
Semantic analysis completed.
`
}
