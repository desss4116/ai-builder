const fs = require("fs")
const path = require("path")

function make(dir) {
  if (!fs.existsSync(dir)) {
    fs.mkdirSync(dir, { recursive: true })
  }
}

function save(file, content) {
  make(path.dirname(file))
  fs.writeFileSync(file, content)
  console.log("CREATED:", file)
}

const structure = [
  "app",
  "app/components",
  "app/components/ui",
  "app/components/layout",
  "app/components/game",
  "app/engine",
  "app/games",
  "app/games/glade-run",
  "app/games/map-maker",
  "app/store",
  "app/hooks",
  "app/lib",
  "app/styles",
  "public"
]

structure.forEach(make)

save(
  "package.json",
  JSON.stringify(
    {
      name: "ai-builder",
      version: "1.0.0",
      private: true,

      scripts: {
        dev: "next dev",
        build: "next build",
        start: "next start"
      },

      dependencies: {
        next: "14.2.35",
        react: "^18.2.0",
        "react-dom": "^18.2.0",
        "framer-motion": "^11.0.0",
        zustand: "^4.5.2",
        three: "^0.164.1",
        "@react-three/fiber": "^8.16.6",
        "@react-three/drei": "^9.105.6",
        "pixi.js": "^8.1.1"
      }
    },
    null,
    2
  )
)

save(
  "next.config.js",
  `
/** @type {import('next').NextConfig} */

const nextConfig = {
  reactStrictMode: true
}

module.exports = nextConfig
`
)

save(
  "app/layout.jsx",
  `
export const metadata = {
  title: "AI Builder",
  description: "Universal cinematic website generator"
}

import "./styles/globals.css"

export default function RootLayout({ children }) {
  return (
    <html lang="en">
      <body>{children}</body>
    </html>
  )
}
`
)

save(
  "app/styles/globals.css",
  `
*{
  margin:0;
  padding:0;
  box-sizing:border-box;
}

html,body{
  background:#050505;
  color:white;
  font-family:Arial;
  overflow-x:hidden;
}

body{
  min-height:100vh;
}

canvas{
  touch-action:none;
}
`
)

save(
  "app/store/useBuilderStore.js",
  `
import { create } from "zustand"

export const useBuilderStore = create((set) => ({
  glitch:false,
  setGlitch:(v)=>set({glitch:v})
}))
`
)

save(
  "app/components/layout/Hero.jsx",
  `
"use client"

import { motion } from "framer-motion"

export default function Hero() {
  return (
    <section
      style={{
        height:"100vh",
        display:"flex",
        justifyContent:"center",
        alignItems:"center",
        flexDirection:"column",
        position:"relative"
      }}
    >
      <motion.h1
        initial={{opacity:0,y:100}}
        animate={{opacity:1,y:0}}
        transition={{duration:1}}
        style={{
          fontSize:"6rem",
          fontWeight:"900",
          letterSpacing:"6px"
        }}
      >
        AI BUILDER
      </motion.h1>

      <motion.p
        initial={{opacity:0}}
        animate={{opacity:1}}
        transition={{delay:1}}
        style={{
          marginTop:"20px",
          opacity:0.7,
          fontSize:"1.2rem"
        }}
      >
        Universal cinematic website generation engine
      </motion.p>
    </section>
  )
}
`
)

save(
  "app/components/game/GlitchOverlay.jsx",
  `
"use client"

import { useBuilderStore } from "../../store/useBuilderStore"

export default function GlitchOverlay(){

  const glitch = useBuilderStore(s=>s.glitch)

  if(!glitch) return null

  return(
    <div
      style={{
        position:"fixed",
        inset:0,
        background:
          "repeating-linear-gradient(0deg,#ff000022 0px,#000 2px,#000 4px)",
        mixBlendMode:"screen",
        pointerEvents:"none",
        zIndex:9999,
        animation:"flicker .1s infinite"
      }}
    />
  )
}
`
)

save(
  "app/games/glade-run/Game.jsx",
  `
"use client"

import { useState } from "react"
import { useBuilderStore } from "../../store/useBuilderStore"

export default function Game(){

  const [dead,setDead] = useState(false)

  const setGlitch = useBuilderStore(s=>s.setGlitch)

  function lose(){

    setDead(true)

    setGlitch(true)

    setTimeout(()=>{
      setGlitch(false)
      setDead(false)
    },3000)
  }

  return(
    <div
      style={{
        marginTop:"60px",
        display:"flex",
        flexDirection:"column",
        alignItems:"center"
      }}
    >
      <button
        onClick={lose}
        style={{
          padding:"20px 40px",
          background:"#111",
          color:"white",
          border:"1px solid #333",
          cursor:"pointer",
          fontSize:"1rem"
        }}
      >
        Simulate Death
      </button>

      {dead && (
        <h2 style={{marginTop:"30px"}}>
          PLAYER TERMINATED
        </h2>
      )}
    </div>
  )
}
`
)

save(
  "app/page.jsx",
  `
"use client"

import Hero from "./components/layout/Hero"
import Game from "./games/glade-run/Game"
import GlitchOverlay from "./components/game/GlitchOverlay"

export default function Home(){

  return(
    <>
      <GlitchOverlay />

      <main>
        <Hero />
        <Game />
      </main>
    </>
  )
}
`
)

console.log("")
console.log("AI BUILDER CREATED")
console.log("READY FOR CLOUDFLARE")
console.log("")
