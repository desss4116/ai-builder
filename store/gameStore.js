
"use client"

import {create} from "zustand"

const useGameStore = create((set)=>({

  energy:100,

  level:1,

  mode:"DIRECTOR",

  glitch:false,

  discoveredSecrets:[],

  setEnergy:(energy)=>set({energy}),

  setLevel:(level)=>set({level}),

  setMode:(mode)=>set({mode}),

  triggerGlitch:()=>{

    set({glitch:true})

    setTimeout(()=>{

      set({glitch:false})

    },500)
  },

  unlockSecret:(secret)=>set(state=>({

    discoveredSecrets:[
      ...state.discoveredSecrets,
      secret
    ]
  }))
}))

export default useGameStore
