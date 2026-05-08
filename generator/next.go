package generator

func GenerateNextProject(prompt string) string {

	return `
generated/

app/
components/
sections/
styles/

package.json
tailwind.config.js
next.config.js
`
}
