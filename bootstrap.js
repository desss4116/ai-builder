// bootstrap.js

const fs = require("fs")
const path = require("path")

function createFile(filePath, content = "") {

  fs.mkdirSync(path.dirname(filePath), {
    recursive: true
  })

  fs.writeFileSync(filePath, content)

  console.log("CREATED:", filePath)

}

//
// ROOT
//

createFile(
  "package.json",
  `{
  "name": "ai-builder",
  "private": true,
  "scripts": {
    "dev": "next dev",
    "build": "next build",
    "start": "next start"
  },
  "dependencies": {
    "next": "^14.2.3",
    "react": "^18.3.1",
    "react-dom": "^18.3.1",
    "zustand": "^4.5.2",
    "framer-motion": "^11.2.10",
    "three": "^0.164.1",
    "@react-three/fiber": "^8.16.6",
    "@react-three/drei": "^9.105.6",
    "discord.js": "^14.15.3",
"typescript": "^5.5.4",
"@types/react": "^18.3.3",
"@types/node": "^20.14.10"
}`
)

createFile(
  "next.config.js",
  `module.exports = {
  reactStrictMode: true
}`
)

createFile(
  "tailwind.config.js",
  `module.exports = {
  content: [
    "./app/**/*.{js,ts,jsx,tsx}",
    "./components/**/*.{js,ts,jsx,tsx}"
  ],
  theme: {
    extend: {}
  },
  plugins: []
}`
)

createFile(
  "tsconfig.json",
  `{
  "compilerOptions": {
    "target": "ES6",
    "lib": ["dom", "esnext"],
    "allowJs": true,
    "skipLibCheck": true,
    "strict": false,
    "module": "esnext",
    "moduleResolution": "bundler",
    "jsx": "preserve"
  }
}`
)

//
// APP
//

createFile(
  "app/layout.tsx",
  `
export default function RootLayout({
  children
}) {

  return (
    <html>
      <body>
        {children}
      </body>
    </html>
  )

}
`
)

createFile(
  "app/page.tsx",
  `
import HeroScene from "../components/cinematic/HeroScene"

export default function Home() {

  return (

    <main className="bg-black text-white min-h-screen">

      <HeroScene />

      <div className="p-20">

        <h1 className="text-6xl font-black">
          AI BUILDER
        </h1>

        <p className="mt-6 text-zinc-400 text-xl">
          Universal Autonomous Experience Generator
        </p>

      </div>

    </main>

  )

}
`
)

createFile(
  "app/globals.css",
  `
@tailwind base;
@tailwind components;
@tailwind utilities;

body {
  background: black;
  color: white;
  overflow-x: hidden;
}
`
)

//
// COMPONENTS
//

createFile(
  "components/cinematic/HeroScene.tsx",
  `
"use client"

import { motion } from "framer-motion"

export default function HeroScene() {

  return (

    <section className="h-screen flex items-center justify-center">

      <motion.div

        animate={{
          opacity: [0.4, 1, 0.4]
        }}

        transition={{
          repeat: Infinity,
          duration: 4
        }}

      >

        <h1 className="text-8xl font-black">
          AI BUILDER
        </h1>

      </motion.div>

    </section>

  )

}
`
)

createFile(
  "components/cinematic/GlitchOverlay.tsx",
  `
export default function GlitchOverlay() {

  return (
    <div>
      Glitch Overlay
    </div>
  )

}
`
)

createFile(
  "components/ui/Button.tsx",
  `
export default function Button({
  children
}) {

  return (

    <button className="px-6 py-3 rounded-xl bg-white text-black">

      {children}

    </button>

  )

}
`
)

//
// SYSTEMS
//

createFile(
  "systems/semanticEngine.ts",
  `
export function analyzeIntent(prompt) {

  const lower = prompt.toLowerCase()

  if(lower.includes("dashboard")) {
    return "dashboard"
  }

  if(lower.includes("game")) {
    return "game"
  }

  if(lower.includes("ai")) {
    return "ai-app"
  }

  return "cinematic"

}
`
)

createFile(
  "systems/websiteGenerator.ts",
  `
import { analyzeIntent } from "./semanticEngine"

export async function generateWebsite(prompt) {

  const type = analyzeIntent(prompt)

  return {

    type,

    pages: [
      "home",
      "about",
      "dashboard"
    ],

    animations: true,

    cinematic: true

  }

}
`
)

createFile(
  "systems/deployEngine.ts",
  `
export async function deployProject(name) {

  return {

    success: true,

    url: "https://example.pages.dev/" + name

  }

}
`
)

createFile(
  "systems/discordEngine.ts",
  `
export async function processDiscordPrompt(prompt) {

  return {
    success: true
  }

}
`
)

//
// STORE
//

createFile(
  "store/appStore.ts",
  `
import { create } from "zustand"

export const useAppStore = create((set) => ({

  theme: "dark",

  setTheme: (theme) => set({
    theme
  })

}))
`
)

//
// TEMPLATES
//

createFile(
  "templates/saas.ts",
  `
export const saasTemplate = {

  hero: true,

  pricing: true,

  dashboard: true

}
`
)

createFile(
  "templates/cinematic.ts",
  `
export const cinematicTemplate = {

  particles: true,

  shaders: true,

  animations: true

}
`
)

//
// DISCORD BOT
//

createFile(
  "discord-bot.js",
  `
const {
  Client,
  GatewayIntentBits
} = require("discord.js")

const client = new Client({

  intents: [
    GatewayIntentBits.Guilds,
    GatewayIntentBits.GuildMessages,
    GatewayIntentBits.MessageContent
  ]

})

client.on("messageCreate", async (message) => {

  if(message.author.bot) return

  const prompt = message.content

  message.reply(
    "Generating website for: " + prompt
  )

})

client.login(process.env.DISCORD_TOKEN)
`
)

console.log("\\nAI BUILDER CREATED SUCCESSFULLY")
