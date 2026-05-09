
import {create} from "zustand"

const useBuilderStore = create((set)=>({

  prompt:"",
  generated:"",

  setPrompt:(prompt)=>set({prompt}),
  setGenerated:(generated)=>set({generated})

}))

export default useBuilderStore
