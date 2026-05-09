
export class ProjectManager{

  constructor(){

    this.projects = []
  }

  create(project){

    this.projects.push(project)

    return project
  }

  getAll(){

    return this.projects
  }
}
