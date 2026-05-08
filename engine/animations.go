package engine

func Animations() string {
	return `
<style>

.fade-up{
	animation:fadeUp 1s ease;
}

@keyframes fadeUp{
	from{
		opacity:0;
		transform:translateY(30px);
	}
	to{
		opacity:1;
		transform:translateY(0);
	}
}

</style>
`
}
