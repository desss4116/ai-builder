
import {create}
from "zustand"

const useEditorStore =

create((set)=>({

  files:[],

  selectedFile:null,

  setFiles:(files)=>
    set({
      files
    }),

  selectFile:(selectedFile)=>
    set({
      selectedFile
    })

}))

export default useEditorStore
