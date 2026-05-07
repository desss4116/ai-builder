package components

import (
	"ai-builder/engine" // Подключаем твою структуру Theme
	"fmt"
)

func RenderHero(theme engine.Theme, title, content string) string {
	return fmt.Sprintf(`
	<section class="relative min-h-screen flex items-center justify-center overflow-hidden">
		<div class="absolute inset-0 opacity-20" style="background-image: url('https://transparenttextures.com');"></div>
		<div class="container mx-auto px-6 z-10 text-center">
			<h1 class="text-7xl md:text-9xl font-black mb-8 leading-none %s" data-aos="zoom-out-up">
				%s
			</h1>
			<p class="text-xl md:text-2xl max-w-3xl mx-auto opacity-70 mb-12" data-aos="fade-up" data-aos-delay="200">
				%s
			</p>
			<div class="flex flex-wrap justify-center gap-6" data-aos="fade-up" data-aos-delay="400">
				<button class="px-12 py-5 text-lg font-bold uppercase tracking-widest %s">Launch Now</button>
				<button class="px-12 py-5 text-lg font-bold uppercase tracking-widest border border-current hover:bg-white hover:text-black transition">Explore</button>
			</div>
		</div>
	</section>`, theme.H1Class, title, content, theme.BtnClass)
}

