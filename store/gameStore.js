
import {create}
from "zustand"

const useGameStore =

create((set)=>({

  health:100,

  stamina:100,

  inventory:[],

  setHealth:(health)=>
    set({
      health
    }),

  setStamina:(stamina)=>
    set({
      stamina
    })

}))

export default useGameStore
