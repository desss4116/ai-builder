package analysis

import "strings"

func DetectStyle(prompt string) string {
	prompt = strings.ToLower(prompt)

	if strings.Contains(prompt, "luxury") {
		return "luxury"
	}

	if strings.Contains(prompt, "minimal") {
		return "minimal"
	}

	if strings.Contains(prompt, "futuristic") {
		return "futuristic"
	}

	return "modern"
}
