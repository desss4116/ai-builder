
const memory = []

export function saveMemory(
  item
){

  memory.push({

    ...item,

    createdAt:
    Date.now()
  })

  return true
}

export function getMemory(){

  return memory
}
