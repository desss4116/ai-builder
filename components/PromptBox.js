
"use client"

import {useState} from "react"
import useBuilderStore from "../store/builderStore"
import {generateWebsite} from "../engine/generation/generator"

export default function PromptBox(){

  const [loading,setLoading] = useState(false)

  const prompt = useBuilderStore(s=>s.prompt)
  const setPrompt = useBuilderStore(s=>s.setPrompt)
  const setGenerated = useBuilderStore(s=>s.setGenerated)

  async function handleGenerate(){

    setLoading(true)

    const result = await generateWebsite(prompt)

    setGenerated(result)

    setLoading(false)
  }

  return (
    <section
      style={{
        maxWidth:"1000px",
        margin:"auto",
        padding:"20px"
      }}
    >

      <textarea
        value={prompt}
        onChange={(e)=>setPrompt(e.target.value)}
        placeholder="Describe your website..."
        style={{
          width:"100%",
          height:"220px",
          background:"#111827",
          color:"white",
          border:"none",
          borderRadius:"20px",
          padding:"20px",
          fontSize:"18px"
        }}
      />

      <button
        onClick={handleGenerate}
        style={{
          width:"100%",
          marginTop:"20px",
          padding:"20px",
          borderRadius:"20px",
          border:"none",
          background:"#2563eb",
          color:"white",
          fontSize:"20px"
        }}
      >
        {loading ? "Generating..." : "Generate Website"}
      </button>

    </section>
  )
}
