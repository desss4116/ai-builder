
"use client"

export default function CinematicOverlay(){

  return (

    <div

      style={{

        position:"fixed",

        inset:0,

        pointerEvents:"none",

        background:
        "linear-gradient(to bottom, transparent, rgba(0,0,0,0.4))"
      }}
    />
  )
}
