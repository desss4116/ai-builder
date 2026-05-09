
"use client"

import useBuilderStore
from "../store/builderStore"

export default function Preview(){

  const generated =
  useBuilderStore(s=>s.generated)

  return (
    <section
      style={{
        padding:"40px"
      }}
    >

      <div
        className="glass"
        style={{
          borderRadius:"24px",
          padding:"30px",
          minHeight:"500px"
        }}
      >

        <h2
          style={{
            marginBottom:"20px",
            fontSize:"32px"
          }}
        >
          AI Output
        </h2>

        <pre
          style={{
            whiteSpace:"pre-wrap",
            lineHeight:"1.8",
            opacity:0.85
          }}
        >
          {generated}
        </pre>

      </div>

    </section>
  )
}
