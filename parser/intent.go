package parser

import "strings"

func IsWebsiteRequest(text string) bool {

	text = strings.ToLower(text)

	keywords := []string{
		"создай",
		"сделай",
		"website",
		"site",
		"landing",
		"лендинг",
		"магазин",
		"сайт",
		"portfolio",
		"startup",
	}

	for _, k := range keywords {
		if strings.Contains(text, k) {
			return true
		}
	}

	return false
}
