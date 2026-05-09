
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
