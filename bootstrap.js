const fs = require("fs")

/*
========================================
FOLDERS
========================================
*/

const folders = [
  "app",
  "app/api",
  "app/api/generate",
  "components",
  "engine",
  "games",
  "systems",
  "store",
  "hooks",
  "styles",
  "public",
  "lib"
]

folders.forEach(folder => {
  fs.mkdirSync(folder, {
    recursive: true
  })
})

/*
========================================
PACKAGE.JSON
========================================
*/

fs.writeFileSync(
  "package.json",
  JSON.stringify(
    {
      name: "ai-builder",
      private: true,
      version: "1.0.0",
      scripts: {
        dev: "next dev",
        build: "next build",
        start: "next start"
      },
      dependencies: {
        next: "14.2.35",
        react: "^18.2.0",
        "react-dom": "^18.2.0"
      }
    },
    null,
    2
  )
)

/*
========================================
NEXT CONFIG
========================================
*/

fs.writeFileSync(
  "next.config.js",
`
/** @type {import('next').NextConfig} */

const nextConfig = {}

module.exports = nextConfig
`
)

/*
========================================
GLOBAL CSS
========================================
*/

fs.writeFileSync(
  "app/globals.css",
`
*{
  margin:0;
  padding:0;
  box-sizing:border-box;
}

html,
body{
  background:#050816;
  color:white;
  font-family:Arial;
}
`
)

/*
========================================
ROOT LAYOUT
========================================
*/

fs.writeFileSync(
  "app/layout.js",
`
import "./globals.css"

export const metadata = {
  title:"AI Builder",
  description:"Autonomous Website Intelligence System"
}

export default function RootLayout({ children }) {

  return (
    <html lang="en">
      <body>
        {children}
      </body>
    </html>
  )
}
`
)

/*
========================================
NAVBAR
========================================
*/

fs.writeFileSync(
  "components/Navbar.js",
`
export default function Navbar() {

  return (
    <nav
      style={{
        width:"100%",
        padding:"20px 40px",
        display:"flex",
        justifyContent:"space-between",
        alignItems:"center",
        borderBottom:"1px solid rgba(255,255,255,0.1)"
      }}
    >

      <h2>AI Builder</h2>

      <button
        style={{
          background:"#2563eb",
          color:"white",
          border:"none",
          padding:"12px 20px",
          borderRadius:"10px",
          cursor:"pointer"
        }}
      >
        Dashboard
      </button>

    </nav>
  )
}
`
)

/*
========================================
HERO
========================================
*/

fs.writeFileSync(
  "components/Hero.js",
`
export default function Hero() {

  return (
    <section
      style={{
        minHeight:"90vh",
        display:"flex",
        justifyContent:"center",
        alignItems:"center",
        flexDirection:"column",
        textAlign:"center",
        padding:"20px"
      }}
    >

      <h1
        style={{
          fontSize:"70px",
          marginBottom:"20px"
        }}
      >
        AI Builder
      </h1>

      <p
        style={{
          opacity:0.7,
          maxWidth:"700px",
          lineHeight:"1.6"
        }}
      >
        Autonomous Website Intelligence System
      </p>

    </section>
  )
}
`
)

/*
========================================
HOME PAGE
========================================
*/

fs.writeFileSync(
  "app/page.js",
`
import Navbar from "../components/Navbar"
import Hero from "../components/Hero"

export default function Home() {

  return (
    <main>

      <Navbar />

      <Hero />

    </main>
  )
}
`
)

/*
========================================
API GENERATOR
========================================
*/

fs.writeFileSync(
  "app/api/generate/route.js",
`
export async function POST(req) {

  const body = await req.json()

  return Response.json({
    success:true,
    prompt:body.prompt || null,
    message:"AI Builder Generator Online"
  })
}
`
)

/*
========================================
AI ENGINE
========================================
*/

fs.writeFileSync(
  "engine/brain.js",
`
export function analyzePrompt(prompt) {

  return {
    prompt,
    type:"website",
    style:"modern",
    status:"analyzed"
  }
}
`
)

/*
========================================
SYSTEMS
========================================
*/

fs.writeFileSync(
  "systems/generator.js",
`
export function generateWebsite(data) {

  return {
    success:true,
    website:data
  }
}
`
)

/*
========================================
STORE
========================================
*/

fs.writeFileSync(
  "store/state.js",
`
export const globalState = {
  projects:[]
}
`
)

/*
========================================
HOOKS
========================================
*/

fs.writeFileSync(
  "hooks/useAI.js",
`
export function useAI() {

  async function generate(prompt) {

    const res = await fetch("/api/generate",{
      method:"POST",
      headers:{
        "Content-Type":"application/json"
      },
      body:JSON.stringify({
        prompt
      })
    })

    return res.json()
  }

  return {
    generate
  }
}
`
)

/*
========================================
README
========================================
*/

fs.writeFileSync(
  "README.md",
`
# AI Builder

Autonomous Website Intelligence System
`
)

console.log("AI Builder Full Structure Generated")
