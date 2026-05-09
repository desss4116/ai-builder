
import "./globals.css"

export const metadata = {

  title:"WCKD OS",

  description:"Maze Runner Experience"
}

export default function RootLayout({children}){

  return (

    <html lang="en">

      <body>

        {children}

      </body>

    </html>
  )
}
