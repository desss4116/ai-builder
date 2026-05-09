
"use client"

export default function HologramButton({
  children
}){

  return (

    <button

      style={{

        padding:"20px 40px",

        borderRadius:"20px",

        border:
        "1px solid rgba(255,255,255,0.1)",

        background:
        "rgba(255,255,255,0.05)",

        color:"white",

        backdropFilter:
        "blur(20px)"
      }}
    >

      {children}

    </button>
  )
}
