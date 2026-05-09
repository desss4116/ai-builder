
export class AIDirector{

  constructor(){

    this.state = "idle"
  }

  decide(context){

    if(
      context.danger
    ){

      return {
        action:"escape"
      }
    }

    return {
      action:"explore"
    }
  }
}
