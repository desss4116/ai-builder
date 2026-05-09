
export function saveGame(
  data
){

  localStorage.setItem(

    "ai-builder-save",

    JSON.stringify(data)
  )

  return true
}

export function loadGame(){

  const data =

  localStorage.getItem(
    "ai-builder-save"
  )

  return JSON.parse(data)
}
