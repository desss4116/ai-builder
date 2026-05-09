
export class Multiplayer{

  constructor(){

    this.players = []
  }

  connect(player){

    this.players.push(player)
  }

  broadcast(message){

    return {

      success:true,

      players:
      this.players.length,

      message
    }
  }
}
