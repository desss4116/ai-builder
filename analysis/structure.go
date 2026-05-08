package analysis

import "strings"

type StructureAnalysis struct {
	HasNavbar  bool
	HasHero    bool
	HasCards   bool
	HasFooter  bool
	HasPricing bool
}

func AnalyzeStructure(html string) StructureAnalysis {
	html = strings.ToLower(html)

	return StructureAnalysis{
		HasNavbar:  strings.Contains(html, "<nav"),
		HasHero:    strings.Contains(html, "hero"),
		HasCards:   strings.Contains(html, "card"),
		HasFooter:  strings.Contains(html, "<footer"),
		HasPricing: strings.Contains(html, "pricing"),
	}
}
