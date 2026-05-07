package engine

import "strings"

type PromptContext struct {
	Type     string
	Business string
	Style    string
}

func ParsePrompt(prompt string) PromptContext {

	p := strings.ToLower(prompt)

	ctx := PromptContext{
		Type:  "landing",
		Style: "modern",
	}

	if strings.Contains(p, "крос") ||
		strings.Contains(p, "shoe") ||
		strings.Contains(p, "sneaker") {

		ctx.Business = "sneakers"
		ctx.Type = "ecommerce"
	}

	if strings.Contains(p, "coffee") ||
		strings.Contains(p, "кофе") {

		ctx.Business = "coffee"
		ctx.Type = "restaurant"
	}

	if strings.Contains(p, "portfolio") {

		ctx.Business = "portfolio"
		ctx.Type = "portfolio"
	}

	return ctx
}
