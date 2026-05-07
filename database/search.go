package database

import "fmt"

type KnowledgeResult struct {
	Content string
	Source  string
}

func GetKnowledge(tag string) []KnowledgeResult {
	rows, err := DB.Query("SELECT content, source FROM knowledge WHERE tag LIKE ?", "%"+tag+"%")
	if err != nil {
		return nil
	}
	defer rows.Close()

	var results []KnowledgeResult
	for rows.Next() {
		var r KnowledgeResult
		rows.Scan(&r.Content, &r.Source)
		results = append(results, r)
	}
	return results
}
