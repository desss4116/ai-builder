package engine

import "fmt"

func GenerateWebsite(ctx PromptContext) string {

	theme := RandomTheme()

	title := "Modern Website"

	if ctx.Type == "ecommerce" {
		title = "Sneaker Store"
	}

	if ctx.Type == "restaurant" {
		title = "Coffee Experience"
	}

	if ctx.Type == "portfolio" {
		title = "Creative Portfolio"
	}

	navbar := Navbar()

	hero := Hero(title, theme.Accent)

	features := Features(theme.Card)

	footer := Footer()

	return fmt.Sprintf(`
<!DOCTYPE html>

<html>

<head>

<meta charset="UTF-8">

<meta name="viewport" content="width=device-width, initial-scale=1.0">

<title>%s</title>

<style>

body{
margin:0;
font-family:Arial;
background:%s;
color:white;
}

a{
text-decoration:none;
}

</style>

</head>

<body>

%s

%s

%s

%s

</body>

</html>
`,
		title,
		theme.Background,
		navbar,
		hero,
		features,
		footer,
	)
}
