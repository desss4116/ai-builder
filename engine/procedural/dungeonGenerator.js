
export function generateDungeon(){

  const rooms = []

  for(let i=0;i<20;i++){

    rooms.push({

      id:i,

      x:
      Math.random() * 100,

      y:
      Math.random() * 100
    })
  }

  return rooms
}
