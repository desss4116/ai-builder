
"use client"

import {useEffect}
from "react"

export default function PixiGame(){

  useEffect(()=>{

    console.log(
      "Pixi Runtime Online"
    )

  },[])

  return (

    <div
      style={{
        padding:"40px",
        background:"#111827",
        borderRadius:"24px"
      }}
    >

      PixiJS Game Runtime

    </div>
  )
}
