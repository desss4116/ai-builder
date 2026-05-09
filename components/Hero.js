
"use client"

import {motion}
from "framer-motion"

export default function Hero(){

  return (

    <section

      style={{
        padding:"120px 20px",
        textAlign:"center"
      }}
    >

      <motion.h1

        initial={{
          opacity:0,
          y:40
        }}

        animate={{
          opacity:1,
          y:0
        }}

        transition={{
          duration:1
        }}

        style={{
          fontSize:"92px",
          marginBottom:"20px"
        }}
      >

        AI Builder Omega Ultra

      </motion.h1>

      <motion.p

        initial={{
          opacity:0
        }}

        animate={{
          opacity:1
        }}

        transition={{
          delay:0.4
        }}

        style={{
          fontSize:"24px",
          opacity:0.7
        }}
      >

        Autonomous Website + Game AI Factory

      </motion.p>

    </section>
  )
}
