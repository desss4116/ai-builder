
import "./globals.css"

export const metadata = {

  title:"AI Builder Omega Ultra",

  description:
  "Autonomous Website + Game AI Factory"
}

export default function RootLayout({
  children
}){

  return (

    <html lang="en">

      <body>

        {children}

      </body>

    </html>
  )
}
