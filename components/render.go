package components

import "fmt"

func RenderButton(text string, link string) string {
	return fmt.Sprintf(`
<a href="%s" class="btn">
	%s
</a>
`, link, text)
}

func RenderCard(title string, desc string) string {
	return fmt.Sprintf(`
<div class="card">
	<h3>%s</h3>
	<p>%s</p>
</div>
`, title, desc)
}
