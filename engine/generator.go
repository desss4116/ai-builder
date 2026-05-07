package engine

import (
	"ai-builder/components"
	"fmt"
)

func GenerateWebsite(prompt string) string {
	theme := GetTheme(prompt)
	context := SearchInternet(prompt)

	return fmt.Sprintf(`
<!DOCTYPE html>
<html lang="ru" class="scroll-smooth">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <script src="https://tailwindcss.com"></script>
    <link rel="stylesheet" href="https://unpkg.com" />
    <style>
        @import url('https://googleapis.com');
        body { font-family: 'Plus Jakarta Sans', sans-serif; background: %s; color: %s; }
        .glass { background: rgba(255,255,255,0.03); backdrop-filter: blur(12px); border: 1px solid rgba(255,255,255,0.05); }
        .text-gradient { background: linear-gradient(to right, #fff, %s); -webkit-background-clip: text; -webkit-text-fill-color: transparent; }
    </style>
</head>
<body class="antialiased overflow-x-hidden">
    %s %s %s %s
    <script src="https://unpkg.com"></script>
    <script> AOS.init({ duration: 1000, once: true }); </script>
</body></html>`, theme.Bg, theme.Text, theme.Accent, components.RenderNavbar(theme), components.RenderHero(theme, prompt, context), components.RenderFeatures(theme), components.RenderFooter(theme))
}
