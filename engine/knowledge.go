package engine

type BusinessKnowledge struct {
	Sections []string
}

var EcommerceKnowledge = BusinessKnowledge{
	Sections: []string{
		"hero",
		"products",
		"reviews",
		"cta",
	},
}

var RestaurantKnowledge = BusinessKnowledge{
	Sections: []string{
		"hero",
		"menu",
		"gallery",
		"reviews",
	},
}
