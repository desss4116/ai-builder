
"use client"

export default function CodeEditor(){

  return (

    <textarea

      style={{

        width:"100%",

        height:"600px",

        background:"#020617",

        color:"#22c55e",

        border:"none",

        borderRadius:"24px",

        padding:"24px",

        fontFamily:"monospace"
      }}

      defaultValue={`

// AI Builder Runtime

export default function App(){

  return <div>Hello</div>
}

      `}
    />
  )
}
