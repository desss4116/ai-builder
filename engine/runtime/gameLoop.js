
export class GameLoop{

  start(){

    requestAnimationFrame(
      this.update.bind(this)
    )
  }

  update(){

    requestAnimationFrame(
      this.update.bind(this)
    )
  }
}
