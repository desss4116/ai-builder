
export function synthesizeTSX(
  componentName
){

  return `

export default function ${componentName}(){

  return (

    <section>

      <h1>

        ${componentName}

      </h1>

    </section>
  )
}

  `
}
