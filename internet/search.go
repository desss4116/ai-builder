package internet

import "strings"

func SearchWeb(query string) string {

	q := strings.ToLower(query)

	if strings.Contains(q, "фильм") {

		return `
Interstellar
Oppenheimer
Dune Part Two
Blade Runner 2049
The Batman
`
	}

	if strings.Contains(q, "apple") {

		return `
Apple is known for:
premium UX
minimalism
ecosystem design
smooth animations
`
	}

	if strings.Contains(q, "netflix") {

		return `
Netflix uses:
dark cinematic UI
personalization
streaming architecture
motion transitions
`
	}

	return `
Global internet analysis completed.
`
}
