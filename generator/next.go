package generator

func GenerateNextProject(prompt string) string {

	return `
generated/

app/
components/
sections/
styles/
public/

app/page.tsx
app/layout.tsx

components/Hero.tsx
components/Navbar.tsx
components/Footer.tsx

styles/globals.css

tailwind.config.js
next.config.js
package.json

Features:
- Next.js
- TailwindCSS
- Framer Motion
- Responsive UI
- Premium Design
- Modern Architecture
`
}
