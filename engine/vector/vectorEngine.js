
export function createVector(
  text
){

  const tokens =
  text.split(" ")

  return {

    dimensions:128,

    tokens,

    length:tokens.length
  }
}
