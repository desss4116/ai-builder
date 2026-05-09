const fs = require("fs")

const folders = [
  "app",
  "components",
  "engine",
  "games",
  "systems",
  "store",
  "hooks",
  "shaders",
  "styles",
  "public",
  "lib"
]

folders.forEach(folder => {
  if (!fs.existsSync(folder)) {
    fs.mkdirSync(folder, { recursive: true })
  }
})

fs.writeFileSync(
  "app/page.js",
`
export default function Home() {
  return (
    <main style={{
      background:"#050816",
      color:"white",
      height:"100vh",
      display:"flex",
      justifyContent:"center",
      alignItems:"center",
      flexDirection:"column",
      fontFamily:"Arial"
    }}>
      <h1 style={{
        fontSize:"60px",
        color:"#00ffff"
      }}>
        AI BUILDER
      </h1>

      <p>
        Autonomous Website Intelligence System
      </p>
    </main>
  )
}
`
)

fs.writeFileSync(
  "package.json",
`
{
  "name": "ai-builder",
  "private": true,
  "scripts": {
    "build": "next build",
    "dev": "next dev"
  },
  "dependencies": {
    "next": "14.2.35",
    "react": "^18.2.0",
    "react-dom": "^18.2.0"
  }
}
`
)

console.log("AI Builder Generated")
