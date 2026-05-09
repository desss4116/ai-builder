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

const page = `
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

fs.writeFileSync("app/page.js", page)

const packageJson = `
{
  "name": "ai-builder",
  "private": true,
  "scripts": {
    "dev": "next dev",
    "build": "next build"
  },
  "dependencies": {
    "next": "14.2.35",
    "react": "^18.2.0",
    "react-dom": "^18.2.0"
  }
}
`

fs.writeFileSync("package.json", packageJson)

console.log("AI Builder Generated")
