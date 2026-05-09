
"use client"

import {motion} from "framer-motion"

import useGameStore from "../../store/gameStore"

export default function TerminalOverlay(){

  const glitch = useGameStore(s=>s.glitch)

  const energy = useGameStore(s=>s.energy)

  const level = useGameStore(s=>s.level)

  return (

    <motion.div

      animate={{
        opacity:glitch ? 0.4 : 1
      }}

      style={{

        position:"fixed",

        inset:0,

        pointerEvents:"none",

        padding:"20px",

        display:"flex",

        flexDirection:"column",

        justifyContent:"space-between"
      }}
    >

      <div>

        <h1
          className="glitch"
          data-text="WCKD TERMINAL"
          style={{

            fontSize:"42px",

            color:"#00ffee"
          }}
        >
          WCKD TERMINAL
        </h1>

      </div>

      <div>

        <div>ENERGY: {energy}</div>

        <div>LEVEL: {level}</div>

        <div>STATUS: ONLINE</div>

      </div>

    </motion.div>
  )
}
