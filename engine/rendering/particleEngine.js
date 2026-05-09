
export function createParticles(){

  const particles = []

  for(let i=0;i<1000;i++){

    particles.push({

      x:Math.random() * 100,

      y:Math.random() * 100,

      z:Math.random() * 100
    })
  }

  return particles
}
