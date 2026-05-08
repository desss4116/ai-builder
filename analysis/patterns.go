package analysis

import "strings"

func DetectPatterns(html string) []string {
	var patterns []string

	html = strings.ToLower(html)

	if strings.Contains(html, "glass") {
		patterns = append(patterns, "glassmorphism")
	}

	if strings.Contains(html, "gradient") {
		patterns = append(patterns, "gradients")
	}

	if strings.Contains(html, "shadow") {
		patterns = append(patterns, "soft-shadows")
	}

	if strings.Contains(html, "transition") {
		patterns = append(patterns, "animations")
	}

	if strings.Contains(html, "grid") {
		patterns = append(patterns, "grid-layout")
	}

	return patterns
}
