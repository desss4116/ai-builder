package parser

import "strings"

func IsWebsiteRequest(msg string) bool {
	msg = strings.ToLower(msg)

	keywords := []string{
		"создай",
		"site",
		"website",
		"landing",
		"лендинг",
		"сайт",
	}

	for _, k := range keywords {
		if strings.Contains(msg, k) {
			return true
		}
	}

	return false
}
