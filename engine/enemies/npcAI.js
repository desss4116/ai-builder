
export class NPC{

  constructor(name){

    this.name = name

    this.health = 100

    this.state = "idle"
  }

  think(){

    return {

      action:"search",

      target:"player"
    }
  }
}
