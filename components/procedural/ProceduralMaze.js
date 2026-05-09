
"use client"

export default function ProceduralMaze(){

  const cells = []

  for(let i=0;i<400;i++){

    cells.push(

      <div

        key={i}

        style={{

          width:"20px",

          height:"20px",

          background:
          Math.random() > 0.7
          ? "#2563eb"
          : "#020617",

          border:
          "1px solid #111827"
        }}
      />
    )
  }

  return (

    <div

      style={{

        display:"grid",

        gridTemplateColumns:
        "repeat(20,20px)",

        gap:"2px"
      }}
    >

      {cells}

    </div>
  )
}
