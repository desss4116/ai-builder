
export class PhysicsEngine{

  constructor(){

    this.gravity = 9.81

    this.objects = []
  }

  add(object){

    this.objects.push(object)
  }

  update(delta){

    this.objects.forEach((object)=>{

      if(object.velocity){

        object.position.y -=
        this.gravity * delta
      }
    })
  }
}
