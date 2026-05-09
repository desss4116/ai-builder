/*========================================
AI BUILDER OMEGA ULTRA BOOTSTRAP
WCKD OS TERMINAL GENESIS
========================================*/

const fs = require("fs")
const path = require("path")

function ensure(dir){

  if(!fs.existsSync(dir)){

    fs.mkdirSync(
      dir,
      {recursive:true}
    )
  }
}

function write(file,content){

  ensure(path.dirname(file))

  fs.writeFileSync(
    file,
    content
  )

  console.log("CREATED:",file)
}

/*========================================
ROOT STRUCTURE
========================================*/

const folders = [

  "app",
  "app/api",
  "app/api/chat",
  "app/api/generate",
  "app/api/research",

  "components",
  "components/wckd",
  "components/three",
  "components/game",
  "components/ui",

  "engine",
  "engine/ai",
  "engine/runtime",
  "engine/generation",
  "engine/research",
  "engine/deployment",

  "store",

  "public",
  "styles",

  "runtime",
  "runtime/projects",

  "discord",

  "lib"
]

folders.forEach(ensure)

/*========================================
PACKAGE.JSON
========================================*/

write(
  "package.json",
`
{
  "name":"ai-builder-omega-ultra",
  "private":true,
  "version":"1.0.0",

  "scripts":{

    "dev":"next dev",

    "build":"next build",

    "start":"next start"
  },

  "dependencies":{

    "next":"14.2.5",

    "react":"18.2.0",

    "react-dom":"18.2.0",

    "three":"^0.164.1",

    "@react-three/fiber":"^8.16.8",

    "@react-three/drei":"^9.105.6",

    "framer-motion":"^11.2.10",

    "zustand":"^4.5.2",

    "gsap":"^3.12.5",

    "discord.js":"^14.15.3"
  }
}
`
)

/*========================================
NEXT CONFIG
========================================*/

write(
  "next.config.js",
`
/** @type {import('next').NextConfig} */

const nextConfig = {

  reactStrictMode:true
}

module.exports = nextConfig
`
)

/*========================================
GLOBAL CSS
========================================*/

write(
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

  overflow:hidden;
}

body::before{

  content:"";

  position:fixed;

  inset:0;

  pointer-events:none;

  background:
    repeating-linear-gradient(
      to bottom,
      rgba(255,255,255,0.03),
      rgba(255,255,255,0.03) 1px,
      transparent 1px,
      transparent 3px
    );

  mix-blend-mode:overlay;

  opacity:0.25;
}

.glitch{

  position:relative;
}

.glitch::before,
.glitch::after{

  content:attr(data-text);

  position:absolute;

  left:0;

  top:0;

  width:100%;

  overflow:hidden;
}

.glitch::before{

  transform:translateX(-2px);

  color:#00ffee;

  opacity:0.6;
}

.glitch::after{

  transform:translateX(2px);

  color:#ff0055;

  opacity:0.6;
}
`
)

/*========================================
ROOT LAYOUT
========================================*/

write(
  "app/layout.js",
`
import "./globals.css"

export const metadata = {

  title:"WCKD OS",

  description:"Maze Runner Experience"
}

export default function RootLayout({children}){

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

/*========================================
ZUSTAND STORE
========================================*/

write(
  "store/gameStore.js",
`
"use client"

import {create} from "zustand"

const useGameStore = create((set)=>({

  energy:100,

  level:1,

  mode:"DIRECTOR",

  glitch:false,

  discoveredSecrets:[],

  setEnergy:(energy)=>set({energy}),

  setLevel:(level)=>set({level}),

  setMode:(mode)=>set({mode}),

  triggerGlitch:()=>{

    set({glitch:true})

    setTimeout(()=>{

      set({glitch:false})

    },500)
  },

  unlockSecret:(secret)=>set(state=>({

    discoveredSecrets:[
      ...state.discoveredSecrets,
      secret
    ]
  }))
}))

export default useGameStore
`
)

/*========================================
AI ENGINE
========================================*/

write(
  "engine/ai/reasoner.js",
`
export async function reason(query){

  const lower = query.toLowerCase()

  let category = "general"

  if(lower.includes("maze")){

    category = "maze"
  }

  if(lower.includes("game")){

    category = "game"
  }

  if(lower.includes("website")){

    category = "website"
  }

  return {

    success:true,

    category,

    response:
      "WCKD AI analyzed your request and generated tactical response.",

    recommendations:[

      "Enable cinematic mode",

      "Activate procedural systems",

      "Increase immersion layer"
    ]
  }
}
`
)

/*========================================
RESEARCH ENGINE
========================================*/

write(
  "engine/research/search.js",
`
export async function searchInternet(query){

  return {

    success:true,

    query,

    results:[

      {

        title:"WCKD Research Result",

        snippet:
          "Advanced cinematic systems initialized."
      }
    ]
  }
}
`
)

/*========================================
GENERATION ENGINE
========================================*/

write(
  "engine/generation/generator.js",
`
export async function generateProject(prompt){

  return {

    success:true,

    project:"maze-runner",

    deployment:
      "https://maze-runner.pages.dev",

    features:[

      "Three.js Maze",

      "WCKD UI",

      "Glitch Effects",

      "AI Runtime",

      "Director Mode"
    ]
  }
}
`
)

/*========================================
THREE MAZE
========================================*/

write(
  "components/three/MazeScene.js",
`
"use client"

import {Canvas,useFrame} from "@react-three/fiber"

import {OrbitControls} from "@react-three/drei"

import {useRef} from "react"

function MazeWall(props){

  return (

    <mesh {...props}>

      <boxGeometry args={[1,2,1]} />

      <meshStandardMaterial color="#00ffee" />

    </mesh>
  )
}

function Player(){

  const ref = useRef()

  useFrame(()=>{

    if(ref.current){

      ref.current.rotation.y += 0.01
    }
  })

  return (

    <mesh
      ref={ref}
      position={[0,0.5,0]}
    >

      <sphereGeometry args={[0.3,32,32]} />

      <meshStandardMaterial color="#ffffff" />

    </mesh>
  )
}

export default function MazeScene(){

  const walls = []

  for(let x=-5;x<5;x++){

    for(let z=-5;z<5;z++){

      if(Math.random()>0.7){

        walls.push(

          <MazeWall
            key={x+"-"+z}
            position={[x,1,z]}
          />
        )
      }
    }
  }

  return (

    <Canvas camera={{position:[0,8,10]}}>

      <ambientLight intensity={0.5} />

      <pointLight
        position={[10,10,10]}
        intensity={2}
      />

      {walls}

      <Player />

      <OrbitControls />

    </Canvas>
  )
}
`
)

/*========================================
WCKD TERMINAL UI
========================================*/

write(
  "components/wckd/TerminalOverlay.js",
`
"use client"

import {motion} from "framer-motion"

import useGameStore from "../../store/gameStore"

export default function TerminalOverlay(){

  const glitch = useGameStore(s=>s.glitch)

  const energy = useGameStore(s=>s.energy)

  const level = useGameStore(s=>s.level)

  return (

    <motion.div

      animate={{
        opacity:glitch ? 0.4 : 1
      }}

      style={{

        position:"fixed",

        inset:0,

        pointerEvents:"none",

        padding:"20px",

        display:"flex",

        flexDirection:"column",

        justifyContent:"space-between"
      }}
    >

      <div>

        <h1
          className="glitch"
          data-text="WCKD TERMINAL"
          style={{

            fontSize:"42px",

            color:"#00ffee"
          }}
        >
          WCKD TERMINAL
        </h1>

      </div>

      <div>

        <div>ENERGY: {energy}</div>

        <div>LEVEL: {level}</div>

        <div>STATUS: ONLINE</div>

      </div>

    </motion.div>
  )
}
`
)

/*========================================
HOME PAGE
========================================*/

write(
  "app/page.js",
`
"use client"

import dynamic from "next/dynamic"

import TerminalOverlay
from "../components/wckd/TerminalOverlay"

const MazeScene = dynamic(

  ()=>import("../components/three/MazeScene"),

  {ssr:false}
)

export default function Home(){

  return (

    <main
      style={{
        width:"100vw",
        height:"100vh"
      }}
    >

      <MazeScene />

      <TerminalOverlay />

    </main>
  )
}
`
)

/*========================================
CHAT API
========================================*/

write(
  "app/api/chat/route.js",
`
import {reason}
from "../../../engine/ai/reasoner"

export async function POST(req){

  const body = await req.json()

  const result = await reason(body.message)

  return Response.json(result)
}
`
)

/*========================================
GENERATION API
========================================*/

write(
  "app/api/generate/route.js",
`
import {generateProject}
from "../../../engine/generation/generator"

export async function POST(req){

  const body = await req.json()

  const result =
    await generateProject(body.prompt)

  return Response.json(result)
}
`
)

/*========================================
DISCORD BOT
========================================*/

write(
  "discord/bot.js",
`
const {
  Client,
  GatewayIntentBits
} = require("discord.js")

const client = new Client({

  intents:[

    GatewayIntentBits.Guilds,

    GatewayIntentBits.GuildMessages,

    GatewayIntentBits.MessageContent
  ]
})

client.on("ready",()=>{

  console.log("WCKD AI ONLINE")
})

client.on("messageCreate",async(message)=>{

  if(message.author.bot) return

  const query = message.content

  if(query.toLowerCase().includes("create")){

    return message.reply(\`

🧠 WEBSITE GENERATED

PROJECT:
Maze Runner Experience

FEATURES:
• 3D Maze
• WCKD UI
• Procedural Systems
• Director Mode
• Glitch Overlay

DEPLOYED:
https://maze-runner.pages.dev

\`)
  }

  return message.reply(

    "WCKD AI analyzed your request."
  )
})

client.login(process.env.DISCORD_TOKEN)
`
)

/*========================================
README
========================================*/

write(
  "README.md",
`
# WCKD OS TERMINAL GENESIS

AAA CINEMATIC MAZE RUNNER EXPERIENCE

FEATURES:

- THREE.JS
- R3F
- WCKD UI
- GLITCH FX
- ZUSTAND
- FRAMER MOTION
- AI ENGINE
- DISCORD BOT
`
)

console.log(
  "WCKD OS GENERATED SUCCESSFULLY"
)
