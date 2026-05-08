package analysis

import "strings"

func DetectIntent(text string) string {

	text = strings.ToLower(text)

	if strings.Contains(text, "создай") {
		return "website"
	}

	if strings.Contains(text, "top") ||
		strings.Contains(text, "лучшие") {
		return "ranking"
	}

	return "chat"
}
