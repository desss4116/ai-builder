
"use client"

import {Canvas}
from "@react-three/fiber"

import {
  OrbitControls
}
from "@react-three/drei"

function Box(){

  return (

    <mesh rotation={[1,1,0]}>

      <boxGeometry />

      <meshStandardMaterial
        color="#2563eb"
      />

    </mesh>
  )
}

export default function Scene(){

  return (

    <div
      style={{
        height:"600px"
      }}
    >

      <Canvas>

        <ambientLight />

        <pointLight
          position={[10,10,10]}
        />

        <Box />

        <OrbitControls />

      </Canvas>

    </div>
  )
}
