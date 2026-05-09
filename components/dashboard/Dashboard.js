
"use client"

export default function Dashboard(){

  return (

    <section

      style={{

        padding:"40px",

        display:"grid",

        gridTemplateColumns:
        "repeat(3,1fr)",

        gap:"20px"
      }}
    >

      <div
        style={{
          background:"#111827",
          padding:"30px",
          borderRadius:"20px"
        }}
      >
        Projects
      </div>

      <div
        style={{
          background:"#111827",
          padding:"30px",
          borderRadius:"20px"
        }}
      >
        Deployments
      </div>

      <div
        style={{
          background:"#111827",
          padding:"30px",
          borderRadius:"20px"
        }}
      >
        Runtime
      </div>

    </section>
  )
}
