package engine

import "fmt"

func RenderHero(theme Theme, title, desc string) string {
	return fmt.Sprintf(`
	<section class="h-screen flex flex-col justify-center items-center text-center px-4">
		<h1 class="text-6xl md:text-8xl mb-6 %s" data-aos="%s">%s</h1>
		<p class="max-w-2xl text-xl opacity-60 mb-10" data-aos="fade-in">%s</p>
		<button class="px-10 py-4 font-bold text-lg %s" data-aos="fade-up">Explore More</button>
	</section>`, theme.H1Class, theme.AOS, title, desc, theme.BtnClass)
}

func RenderNavbar(theme Theme) string {
	return fmt.Sprintf(`
	<nav class="fixed w-full z-50 p-6 flex justify-between items-center backdrop-blur-md">
		<div class="font-bold text-2xl uppercase tracking-widest">AI.OS</div>
		<div class="space-x-8 hidden md:flex opacity-70">
			<a href="#">Works</a><a href="#">About</a><a href="#">Contact</a>
		</div>
	</nav>`)
}
