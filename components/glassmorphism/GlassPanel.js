
"use client"

export default function GlassPanel({
  children
}){

  return (

    <div

      style={{

        backdropFilter:
        "blur(20px)",

        background:
        "rgba(255,255,255,0.05)",

        border:
        "1px solid rgba(255,255,255,0.1)",

        borderRadius:"24px",

        padding:"24px"
      }}
    >

      {children}

    </div>
  )
}
