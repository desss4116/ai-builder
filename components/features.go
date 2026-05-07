func RenderFeatures(theme engine.Theme) string {
	features := []string{"Ultra Fast", "AI Powered", "Modern UI", "Scalable"}
	cards := ""
	for i, f := range features {
		cards += fmt.Sprintf(`
		<div class="p-10 border border-current opacity-40 hover:opacity-100 transition group" data-aos="fade-up" data-aos-delay="%d">
			<div class="text-4xl mb-4">0%d</div>
			<h3 class="text-2xl font-bold mb-4 uppercase">%s</h3>
			<p class="opacity-60">High-performance intelligence layer integrated into your core business.</p>
		</div>`, i*100, i+1, f)
	}
	return fmt.Sprintf(`<section class="py-24 px-6"><div class="grid md:grid-cols-4 gap-4">%s</div></section>`, cards)
}
