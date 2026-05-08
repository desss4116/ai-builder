package brain

import (
	"strings"
)

func SemanticCleanup(
	text string,
) string {

	lines := strings.Split(
		text,
		".",
	)

	output := ""

	seen := map[string]bool{}

	count := 0

	for _, line := range lines {

		line =
			strings.TrimSpace(line)

		if len(line) < 20 {
			continue
		}

		lower :=
			strings.ToLower(line)

		/*
		   REMOVE DUPLICATES
		*/

		if seen[lower] {
			continue
		}

		seen[lower] = true

		/*
		   REMOVE BAD CONTENT
		*/

		if strings.Contains(lower, "cookie") ||
			strings.Contains(lower, "advertisement") ||
			strings.Contains(lower, "javascript") {
			continue
		}

		output += "• " + line + ".\n\n"

		count++

		if count >= 8 {
			break
		}
	}

	return output
}
