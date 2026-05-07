func RenderScripts() string {
	return `
	<script src="https://unpkg.com"></script>
	<script src="https://jsdelivr.net"></script>
	<script>
		// Анимации
		AOS.init({ duration: 800, once: false, mirror: true });
		
		// Плавный скролл
		const lenis = new Lenis();
		function raf(time) { lenis.raf(time); requestAnimationFrame(raf); }
		requestAnimationFrame(raf);

		// Логика кнопок
		document.querySelectorAll('button').forEach(btn => {
			btn.addEventListener('click', () => {
				btn.style.transform = 'scale(0.95)';
				setTimeout(() => btn.style.transform = 'scale(1)', 100);
			});
		});
	</script>`
}
