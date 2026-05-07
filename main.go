package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

type Site struct {
	ID      string `json:"id"`
	Prompt  string `json:"prompt"`
	HTML    string `json:"html"`
	Created string `json:"created"`
}

var sites = map[string]Site{}

func main() {

	rand.Seed(time.Now().UnixNano())

	token := os.Getenv("DISCORD_TOKEN")

	if token == "" {
		log.Fatal("DISCORD_TOKEN missing")
	}

	dg, err := discordgo.New("Bot " + token)

	if err != nil {
		log.Fatal(err)
	}

	dg.Identify.Intents =
		discordgo.IntentsGuildMessages |
			discordgo.IntentsDirectMessages |
			discordgo.IntentsMessageContent

	dg.AddHandler(messageHandler)

	err = dg.Open()

	if err != nil {
		log.Fatal(err)
	}

	loadSites()

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/site/", siteHandler)
	http.HandleFunc("/api/generate", apiGenerateHandler)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	log.Println("Bot online")
	log.Println("Server running on port", port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.Bot {
		return
	}

	msg := strings.ToLower(m.Content)

	if !isWebsiteRequest(msg) {

		reply := randomChatReply()

		s.ChannelMessageSend(m.ChannelID, reply)

		return
	}

	go func() {

		typingStart(s, m.ChannelID)

		site := generateWebsite(msg)

		saveSite(site)

		domain := os.Getenv("DOMAIN")

		url := fmt.Sprintf(
			"%s/site/%s",
			domain,
			site.ID,
		)

		dm, err := s.UserChannelCreate(m.Author.ID)

		if err == nil {

			s.ChannelMessageSend(
				dm.ID,
				"✅ Ваш сайт готов:\n"+url,
			)

			s.ChannelMessageSend(
				m.ChannelID,
				"📩 Сайт отправлен в личные сообщения",
			)

		} else {

			s.ChannelMessageSend(
				m.ChannelID,
				"✅ Сайт создан:\n"+url,
			)
		}

	}()
}

func isWebsiteRequest(msg string) bool {

	keywords := []string{
		"создай",
		"сделай",
		"site",
		"website",
		"landing",
		"лендинг",
		"магазин",
		"сайт",
		"portfolio",
		"store",
		"shop",
		"ecommerce",
		"build",
	}

	for _, k := range keywords {

		if strings.Contains(msg, k) {
			return true
		}
	}

	return false
}

func randomChatReply() string {

	replies := []string{
		"👋 Напиши: создай сайт для кофейни",
		"🚀 Я умею создавать сайты",
		"💡 Попробуй: website for sneakers",
		"🎨 Могу сделать landing page",
		"⚡ Создаю современные сайты",
		"🔥 Попробуй: сделай сайт для кроссовок",
	}

	return replies[rand.Intn(len(replies))]
}

func detectSiteType(prompt string) string {

	prompt = strings.ToLower(prompt)

	if strings.Contains(prompt, "coffee") ||
		strings.Contains(prompt, "коф") ||
		strings.Contains(prompt, "restaurant") {
		return "restaurant"
	}

	if strings.Contains(prompt, "shoe") ||
		strings.Contains(prompt, "sneaker") ||
		strings.Contains(prompt, "крос") {
		return "ecommerce"
	}

	if strings.Contains(prompt, "portfolio") {
		return "portfolio"
	}

	if strings.Contains(prompt, "agency") {
		return "agency"
	}

	return "saas"
}

func randomStyle() string {

	styles := []string{
		"dark",
		"glass",
		"luxury",
		"minimal",
		"neon",
	}

	return styles[rand.Intn(len(styles))]
}

func generateWebsite(prompt string) Site {

	siteType := detectSiteType(prompt)

	style := randomStyle()

	html := buildHTML(siteType, style, prompt)

	id := randomID()

	return Site{
		ID:      id,
		Prompt:  prompt,
		HTML:    html,
		Created: time.Now().String(),
	}
}

func buildHTML(siteType string, style string, prompt string) string {

	title := strings.Title(prompt)

	bg := "#0f172a"
	card := "#111827"
	accent := "#3b82f6"

	switch style {

	case "glass":
		bg = "linear-gradient(135deg,#0f172a,#1e293b)"
		card = "rgba(255,255,255,0.08)"
		accent = "#06b6d4"

	case "luxury":
		bg = "#0a0a0a"
		card = "#171717"
		accent = "#d4af37"

	case "minimal":
		bg = "#ffffff"
		card = "#f3f4f6"
		accent = "#111827"

	case "neon":
		bg = "#020617"
		card = "#111827"
		accent = "#22d3ee"
	}

	heroText := "Modern AI Generated Website"

	if siteType == "restaurant" {
		heroText = "Premium Restaurant Experience"
	}

	if siteType == "ecommerce" {
		heroText = "Next Generation Store"
	}

	if siteType == "portfolio" {
		heroText = "Creative Portfolio"
	}

	return fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>

<meta charset="UTF-8">

<meta name="viewport" content="width=device-width, initial-scale=1.0">

<title>%s</title>

<style>

*{
margin:0;
padding:0;
box-sizing:border-box;
}

body{
font-family:Arial;
background:%s;
color:white;
overflow-x:hidden;
}

.hero{
padding:120px 20px;
text-align:center;
}

.title{
font-size:72px;
font-weight:bold;
margin-bottom:20px;
}

.subtitle{
font-size:22px;
opacity:.85;
max-width:700px;
margin:auto;
line-height:1.6;
}

.btn{
margin-top:40px;
display:inline-block;
padding:18px 42px;
background:%s;
border-radius:16px;
text-decoration:none;
color:white;
font-weight:bold;
transition:.3s;
box-shadow:0 10px 40px rgba(0,0,0,.3);
}

.btn:hover{
transform:translateY(-5px) scale(1.03);
}

.grid{
display:grid;
grid-template-columns:repeat(auto-fit,minmax(280px,1fr));
gap:24px;
padding:60px 30px;
}

.card{
background:%s;
padding:30px;
border-radius:24px;
backdrop-filter:blur(12px);
box-shadow:0 10px 40px rgba(0,0,0,.25);
transition:.3s;
}

.card:hover{
transform:translateY(-8px);
}

.card h2{
margin-bottom:12px;
}

.features{
padding:80px 30px;
}

.section-title{
font-size:42px;
margin-bottom:40px;
text-align:center;
}

.preview{
padding:100px 30px;
text-align:center;
}

.mockup{
margin:auto;
max-width:900px;
background:%s;
padding:40px;
border-radius:30px;
box-shadow:0 10px 50px rgba(0,0,0,.4);
}

.footer{
padding:60px;
text-align:center;
opacity:.7;
}

@media(max-width:768px){

.title{
font-size:44px;
}

.subtitle{
font-size:18px;
}

}

</style>

</head>

<body>

<section class="hero">

<div class="title">%s</div>

<div class="subtitle">
%s
</div>

<a href="#" class="btn">
Get Started
</a>

</section>

<section class="features">

<div class="section-title">
Features
</div>

<div class="grid">

<div class="card">
<h2>⚡ Fast</h2>
<p>Ultra lightweight architecture.</p>
</div>

<div class="card">
<h2>🎨 Modern</h2>
<p>Premium responsive interface.</p>
</div>

<div class="card">
<h2>🚀 Scalable</h2>
<p>Concurrent multi-user generation.</p>
</div>

<div class="card">
<h2>🧠 Smart</h2>
<p>Pseudo AI deterministic generation.</p>
</div>

</div>

</section>

<section class="preview">

<div class="section-title">
Live Preview
</div>

<div class="mockup">

<h2 style="margin-bottom:20px">
%s
</h2>

<p style="opacity:.8">
Generated individually for user request.
</p>

</div>

</section>

<div class="footer">
Generated by AI Builder
</div>

</body>

</html>
`, title, bg, accent, card, card, title, heroText, title)
}

func randomID() string {

	chars := "abcdefghijklmnopqrstuvwxyz0123456789"

	result := ""

	for i := 0; i < 12; i++ {
		result += string(chars[rand.Intn(len(chars))])
	}

	return result
}

func saveSite(site Site) {

	sites[site.ID] = site

	data, _ := json.MarshalIndent(sites, "", "  ")

	os.WriteFile("sites.json", data, 0644)
}

func loadSites() {

	data, err := os.ReadFile("sites.json")

	if err != nil {
		return
	}

	json.Unmarshal(data, &sites)
}

func typingStart(s *discordgo.Session, channelID string) {

	s.ChannelTyping(channelID)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, `
<html>

<body style="
background:#0f172a;
color:white;
font-family:Arial;
padding:60px;
">

<h1>AI Builder</h1>

<p>
Discord AI Website Builder
</p>

</body>

</html>
`)
}

func siteHandler(w http.ResponseWriter, r *http.Request) {

	id := strings.TrimPrefix(r.URL.Path, "/site/")

	site, ok := sites[id]

	if !ok {

		http.NotFound(w, r)

		return
	}

	w.Header().Set("Content-Type", "text/html")

	fmt.Fprint(w, site.HTML)
}

func apiGenerateHandler(w http.ResponseWriter, r *http.Request) {

	prompt := r.URL.Query().Get("prompt")

	if prompt == "" {
		prompt = "modern website"
	}

	site := generateWebsite(prompt)

	saveSite(site)

	domain := os.Getenv("DOMAIN")

	url := fmt.Sprintf(
		"%s/site/%s",
		domain,
		site.ID,
	)

	json.NewEncoder(w).Encode(map[string]string{
		"url": url,
	})
}
