
"use client"

export default function ParticleField(){

  const particles = []

  for(let i=0;i<200;i++){

    particles.push(

      <div

        key={i}

        style={{

          position:"absolute",

          width:"2px",

          height:"2px",

          borderRadius:"50%",

          background:"white",

          left:
          Math.random() * 100 + "%",

          top:
          Math.random() * 100 + "%"
        }}
      />
    )
  }

  return (

    <div

      style={{

        position:"fixed",

        inset:0,

        overflow:"hidden"
      }}
    >

      {particles}

    </div>
  )
}
