package database

type SearchResult struct {
	Title   string
	Content string
}

func SearchKnowledge(
	query string,
) []SearchResult {

	rows, err := DB.Query(`
	SELECT
	title,
	content
	FROM knowledge_base
	WHERE knowledge_base MATCH ?
	ORDER BY bm25(knowledge_base)
	LIMIT 5
	`, query)

	if err != nil {
		return nil
	}

	defer rows.Close()

	results := []SearchResult{}

	for rows.Next() {

		var r SearchResult

		rows.Scan(
			&r.Title,
			&r.Content,
		)

		results = append(results, r)
	}

	return results
}
