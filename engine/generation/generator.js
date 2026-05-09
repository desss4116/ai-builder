
import {analyzePrompt} from "../reasoning/analyzer"

export async function generateWebsite(prompt){

  const analysis = analyzePrompt(prompt)

  await new Promise(r=>setTimeout(r,1200))

  return `

========================================

AI WEBSITE GENERATED

========================================

PROMPT:
${prompt}

========================================

TYPE:
${analysis.type}

STYLE:
${analysis.style}

========================================

SECTIONS

- Navbar
- Hero
- Features
- Testimonials
- Pricing
- Footer

========================================

AI COMPONENTS

- Dynamic UI
- Responsive Layout
- SEO Structure
- Modern Design
- Smart Content Blocks

========================================

STATUS:
ONLINE

========================================
`
}
