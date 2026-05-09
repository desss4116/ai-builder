
"use client"

export default function HologramCard({

  title

}){

  return (

    <div

      style={{

        padding:"30px",

        borderRadius:"24px",

        background:
        "linear-gradient(to bottom,#0ea5e9,#2563eb)",

        boxShadow:
        "0 0 40px rgba(37,99,235,0.5)"
      }}
    >

      <h2>

        {title}

      </h2>

    </div>
  )
}
