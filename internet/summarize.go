package internet

import "strings"

func BuildAnswer(
	query string,
	results []string,
) string {

	response := "🌐 Ответ:\n\n"

	used := map[string]bool{}

	count := 0

	for _, r := range results {

		r = strings.TrimSpace(r)

		if r == "" {
			continue
		}

		if used[r] {
			continue
		}

		used[r] = true

		response += "• " + r + "\n"

		count++

		if count >= 10 {
			break
		}
	}

	if count == 0 {
		return "Не удалось найти информацию."
	}

	return response
}
