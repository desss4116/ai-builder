package analysis

import "strings"

func DetectStyle(prompt string) string {

	prompt = strings.ToLower(prompt)

	if strings.Contains(prompt, "luxury") {
		return "luxury"
	}

	if strings.Contains(prompt, "nike") {
		return "sport"
	}

	if strings.Contains(prompt, "apple") {
		return "minimal"
	}

	return "modern"
}
