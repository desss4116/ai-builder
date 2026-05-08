package internet

import (
	"strings"
)

type WebsiteAnalysis struct {
	IsModern       bool
	IsDark         bool
	HasGradient    bool
	HasAnimations  bool
	HasGlassmorph  bool
	Category       string
}

func AnalyzeHTML(html string) WebsiteAnalysis {
	html = strings.ToLower(html)

	analysis := WebsiteAnalysis{}

	// Modern UI detection
	if strings.Contains(html, "tailwind") ||
		strings.Contains(html, "framer") ||
		strings.Contains(html, "next") {
		analysis.IsModern = true
	}

	// Dark mode
	if strings.Contains(html, "#000") ||
		strings.Contains(html, "dark") {
		analysis.IsDark = true
	}

	// Gradients
	if strings.Contains(html, "gradient") {
		analysis.HasGradient = true
	}

	// Animations
	if strings.Contains(html, "animation") ||
		strings.Contains(html, "transition") ||
		strings.Contains(html, "aos") {
		analysis.HasAnimations = true
	}

	// Glassmorphism
	if strings.Contains(html, "backdrop-filter") {
		analysis.HasGlassmorph = true
	}

	// Category detection
	if strings.Contains(html, "shop") ||
		strings.Contains(html, "product") {
		analysis.Category = "ecommerce"
	}

	if strings.Contains(html, "startup") ||
		strings.Contains(html, "saas") {
		analysis.Category = "saas"
	}

	if strings.Contains(html, "portfolio") {
		analysis.Category = "portfolio"
	}

	if analysis.Category == "" {
		analysis.Category = "general"
	}

	return analysis
}
