package components

import ("ai-builder/engine"; "fmt")

func RenderHero(theme engine.Theme, title, context string) string {
	return fmt.Sprintf(`
    <section class="min-h-screen flex items-center justify-center relative px-6 overflow-hidden">
        <div class="absolute inset-0 opacity-30" style="background: radial-gradient(circle at 50%% 50%%, %s 0%%, transparent 70%%);"></div>
        <div class="max-w-6xl mx-auto text-center z-10">
            <div class="inline-block px-4 py-1.5 mb-6 text-xs font-bold tracking-widest uppercase glass rounded-full" data-aos="fade-down">AI.OS Generated</div>
            <h1 class="text-7xl md:text-[10rem] font-extrabold tracking-tighter mb-8 text-gradient leading-none" data-aos="zoom-out">
                %s
            </h1>
            <p class="text-xl md:text-2xl opacity-50 mb-12 max-w-2xl mx-auto font-light" data-aos="fade-up" data-aos-delay="200">
                %s
            </p>
            <div class="flex flex-col sm:flex-row gap-4 justify-center items-center" data-aos="fade-up" data-aos-delay="400">
                <button class="px-12 py-5 bg-white text-black font-bold rounded-full hover:scale-105 transition-all shadow-2xl">Начать проект</button>
                <button class="px-12 py-5 glass rounded-full hover:bg-white/10 transition-all">Узнать больше</button>
            </div>
        </div>
    </section>`, theme.Accent, title, context)
}
