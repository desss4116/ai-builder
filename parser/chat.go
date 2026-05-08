package parser

import (
	"strings"
)

func BuildChatResponse(query string, results []string) string {
	if len(results) == 0 {
		return "Я не нашёл информацию."
	}

	var response strings.Builder

	response.WriteString("🌐 Вот что я нашёл:\n\n")

	for i, r := range results {
		if i >= 5 {
			break
		}

		response.WriteString("• " + r + "\n")
	}

	return response.String()
}
