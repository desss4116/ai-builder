package engine

import "strings"

type Theme struct {
	BodyClass string
	H1Class   string
	BtnClass  string
	AOS       string
}

func GetTheme(prompt string) Theme {
	p := strings.ToLower(prompt)
	if strings.Contains(p, "luxury") || strings.Contains(p, "dark") {
		return Theme{
			BodyClass: "bg-neutral-950 text-stone-100",
			H1Class:   "font-serif uppercase tracking-tighter text-amber-500",
			BtnClass:  "bg-amber-500 text-black rounded-none hover:bg-white transition-all",
			AOS:       "fade-up",
		}
	}
	// Дефолтная "чистая" тема
	return Theme{
		BodyClass: "bg-white text-slate-900",
		H1Class:   "font-sans font-black text-blue-600",
		BtnClass:  "bg-blue-600 text-white rounded-full hover:shadow-xl transition-transform",
		AOS:       "zoom-in",
	}
}
