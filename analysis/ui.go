package analysis

import "strings"

func DetectUIStyle(html string) string {
	html = strings.ToLower(html)

	if strings.Contains(html, "luxury") {
		return "luxury"
	}

	if strings.Contains(html, "minimal") {
		return "minimal"
	}

	if strings.Contains(html, "cyber") {
		return "cyberpunk"
	}

	if strings.Contains(html, "startup") {
		return "startup"
	}

	return "modern"
}
