
"use client"

import useBuilderStore from "../store/builderStore"

export default function Preview(){

  const generated = useBuilderStore(s=>s.generated)

  return (
    <section
      style={{
        padding:"40px"
      }}
    >

      <div
        style={{
          background:"#111827",
          borderRadius:"20px",
          padding:"30px",
          minHeight:"400px"
        }}
      >

        <h2
          style={{
            marginBottom:"20px"
          }}
        >
          Live Preview
        </h2>

        <pre
          style={{
            whiteSpace:"pre-wrap",
            lineHeight:"1.6",
            opacity:0.85
          }}
        >
          {generated}
        </pre>

      </div>

    </section>
  )
}
