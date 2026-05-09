
"use client"

import {motion} from "framer-motion"

export default function Hero(){

  return (
    <section
      style={{
        padding:"120px 20px",
        textAlign:"center"
      }}
    >

      <motion.h1
        initial={{opacity:0,y:40}}
        animate={{opacity:1,y:0}}
        transition={{duration:1}}
        className="glitch"
        data-text="AI Builder Omega"
        style={{
          fontSize:"92px",
          marginBottom:"20px"
        }}
      >
        AI Builder Omega
      </motion.h1>

      <motion.p
        initial={{opacity:0}}
        animate={{opacity:1}}
        transition={{delay:0.5}}
        style={{
          fontSize:"24px",
          opacity:0.7
        }}
      >
        Autonomous Website Factory
      </motion.p>

    </section>
  )
}
