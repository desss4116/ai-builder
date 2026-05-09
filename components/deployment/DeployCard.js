
"use client"

export default function DeployCard({
  url
}){

  return (

    <a

      href={url}

      target="_blank"

      style={{

        display:"block",

        padding:"20px",

        background:"#2563eb",

        borderRadius:"20px",

        color:"white",

        textDecoration:"none"
      }}
    >

      Open Deployment

    </a>
  )
}
