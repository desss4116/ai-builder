package engine

func GlobalStyles() string {
	return `
<style>

*{
	margin:0;
	padding:0;
	box-sizing:border-box;
}

body{
	background:#050505;
	color:white;
	font-family:Arial,sans-serif;
	line-height:1.6;
}

.container{
	width:90%;
	max-width:1400px;
	margin:auto;
}

.btn{
	display:inline-block;
	padding:16px 32px;
	border-radius:14px;
	text-decoration:none;
	background:white;
	color:black;
	font-weight:bold;
	transition:0.3s;
}

.btn:hover{
	transform:translateY(-3px);
}

.card{
	background:#111;
	padding:30px;
	border-radius:20px;
	margin-top:20px;
	border:1px solid rgba(255,255,255,0.1);
}

</style>
`
}
