package generator

func GenerateHero(title string, subtitle string) string {

	return `
export default function Hero() {
	return (
		<section className="min-h-screen flex items-center justify-center bg-black text-white">
			<div className="max-w-6xl mx-auto px-6">
				<h1 className="text-7xl font-bold">
					` + title + `
				</h1>

				<p className="text-xl text-zinc-400 mt-6">
					` + subtitle + `
				</p>
			</div>
		</section>
	)
}
`
}
