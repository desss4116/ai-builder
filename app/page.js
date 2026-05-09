
"use client"

import Hero
from "../components/Hero"

import PromptBox
from "../components/PromptBox"

import Preview
from "../components/live-preview/Preview"

export default function Home(){

  return (

    <main>

      <Hero />

      <PromptBox />

      <Preview />

    </main>
  )
}
