package engine

import "fmt"

func GenerateWebsite(prompt string) string {
	theme := GetTheme(prompt)
	context := SearchInternet(prompt)

	html := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
	<script src="https://tailwindcss.com"></script>
	<link rel="stylesheet" href="https://unpkg.com" />
	<style> @import url('https://googleapis.com'); </style>
</head>
<body class="%s antialiased">
	%s
	%s
	<script src="https://unpkg.com"></script>
	<script> AOS.init({ duration: 1000, once: true }); </script>
</body>
</html>`, theme.BodyClass, RenderNavbar(theme), RenderHero(theme, prompt, context))

	return html
}
