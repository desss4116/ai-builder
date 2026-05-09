
"use client"

export default function GlitchText({
  children
}){

  return (

    <h1

      style={{

        textShadow:
        "2px 0 red, -2px 0 blue",

        animation:
        "glitch 1s infinite"
      }}
    >

      {children}

    </h1>
  )
}
