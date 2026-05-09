
"use client"

import {Canvas,useFrame} from "@react-three/fiber"

import {OrbitControls} from "@react-three/drei"

import {useRef} from "react"

function MazeWall(props){

  return (

    <mesh {...props}>

      <boxGeometry args={[1,2,1]} />

      <meshStandardMaterial color="#00ffee" />

    </mesh>
  )
}

function Player(){

  const ref = useRef()

  useFrame(()=>{

    if(ref.current){

      ref.current.rotation.y += 0.01
    }
  })

  return (

    <mesh
      ref={ref}
      position={[0,0.5,0]}
    >

      <sphereGeometry args={[0.3,32,32]} />

      <meshStandardMaterial color="#ffffff" />

    </mesh>
  )
}

export default function MazeScene(){

  const walls = []

  for(let x=-5;x<5;x++){

    for(let z=-5;z<5;z++){

      if(Math.random()>0.7){

        walls.push(

          <MazeWall
            key={x+"-"+z}
            position={[x,1,z]}
          />
        )
      }
    }
  }

  return (

    <Canvas camera={{position:[0,8,10]}}>

      <ambientLight intensity={0.5} />

      <pointLight
        position={[10,10,10]}
        intensity={2}
      />

      {walls}

      <Player />

      <OrbitControls />

    </Canvas>
  )
}
