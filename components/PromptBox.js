
"use client"

import {useState}
from "react"

import useBuilderStore
from "../store/builderStore"

export default function PromptBox(){

  const [loading,setLoading] =
  useState(false)

  const prompt =
  useBuilderStore(
    s=>s.prompt
  )

  const setPrompt =
  useBuilderStore(
    s=>s.setPrompt
  )

  const setGenerated =
  useBuilderStore(
    s=>s.setGenerated
  )

  async function generate(){

    if(!prompt) return

    setLoading(true)

    try{

      const response =
      await fetch(
        "/api/generate",
        {
          method:"POST",

          headers:{
            "Content-Type":
            "application/json"
          },

          body:JSON.stringify({
            prompt
          })
        }
      )

      const data =
      await response.json()

      setGenerated(
        JSON.stringify(
          data,
          null,
          2
        )
      )

    }catch(error){

      setGenerated(
        JSON.stringify(
          {
            success:false,
            error:error.message
          },
          null,
          2
        )
      )
    }

    setLoading(false)
  }

  return (

    <section

      style={{
        maxWidth:"1200px",
        margin:"auto",
        padding:"20px"
      }}
    >

      <textarea

        value={prompt}

        onChange={(e)=>
          setPrompt(
            e.target.value
          )
        }

        placeholder="Describe your cinematic website or game..."

        style={{

          width:"100%",

          height:"240px",

          background:"#111827",

          color:"white",

          border:"none",

          borderRadius:"24px",

          padding:"24px",

          fontSize:"18px"
        }}
      />

      <button

        onClick={generate}

        style={{

          width:"100%",

          marginTop:"20px",

          padding:"22px",

          border:"none",

          borderRadius:"24px",

          background:"#2563eb",

          color:"white",

          fontSize:"22px"
        }}
      >

        {
          loading
          ? "Generating..."
          : "Generate"
        }

      </button>

    </section>
  )
}
