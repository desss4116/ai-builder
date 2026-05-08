package brain

import (
	"strings"
)

func FormatAnswer(
	query string,
	data string,
) string {

	data = strings.TrimSpace(data)

	if data == "" {

		return "❌ Информация не найдена."
	}

	lines := strings.Split(
		data,
		"\n",
	)

	output := "🌐 Ответ:\n\n"

	count := 0

	for _, line := range lines {

		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		/*
		   REMOVE DUPLICATES
		*/

		if strings.Contains(
			strings.ToLower(line),
			"википедия",
		) {

			continue
		}

		output += line + "\n\n"

		count++

		if count >= 5 {
			break
		}
	}

	/*
	   CLEAN FALLBACK
	*/

	if count == 0 {

		output += data
	}

	return output
}
