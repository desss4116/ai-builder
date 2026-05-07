package engine

import "math/rand"

type Theme struct {
	Background string
	Card       string
	Accent     string
}

func RandomTheme() Theme {

	themes := []Theme{

		{
			Background: "#0f172a",
			Card:       "#111827",
			Accent:     "#3b82f6",
		},

		{
			Background: "#020617",
			Card:       "#1e293b",
			Accent:     "#22d3ee",
		},

		{
			Background: "#0a0a0a",
			Card:       "#171717",
			Accent:     "#d4af37",
		},
	}

	return themes[rand.Intn(len(themes))]
}
