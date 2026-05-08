package generator

func GenerateHero(title string) string {

	return `
export default function Hero() {
	return (
		<section className="min-h-screen bg-black text-white flex items-center justify-center">
			<div className="max-w-6xl mx-auto px-6">

				<h1 className="text-8xl font-bold tracking-tight">
					` + title + `
				</h1>

				<p className="text-zinc-400 text-xl mt-6">
					Premium next generation experience.
				</p>

			</div>
		</section>
	)
}
`
}
