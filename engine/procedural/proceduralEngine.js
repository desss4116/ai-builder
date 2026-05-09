
export function generateMaze(){

  const maze = []

  for(let y=0;y<20;y++){

    const row = []

    for(let x=0;x<20;x++){

      row.push(
        Math.random() > 0.7
        ? 1
        : 0
      )
    }

    maze.push(row)
  }

  return maze
}
