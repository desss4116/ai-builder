package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

type Site struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Niche       string `json:"niche"`
	Description string `json:"description"`
	HTML        string `json:"html"`
	URL         string `json:"url"`
	CreatedAt   string `json:"created_at"`
}

var (
	sitesFile = "sites.json"
	domain    = ""
)

func main() {

	rand.Seed(time.Now().UnixNano())

	token := os.Getenv("DISCORD_TOKEN")

	if token == "" {
		log.Fatal("DISCORD_TOKEN not found")
	}

	domain = os.Getenv("DOMAIN")

	if domain == "" {
		log.Fatal("DOMAIN not found")
	}

	dg, err := discordgo.New("Bot " + token)

	if err != nil {
		log.Fatal(err)
	}

	dg.Identify.Intents =
		discordgo.IntentsGuildMessages |
			discordgo.IntentsMessageContent

	dg.AddHandler(onMessage)

	err = dg.Open()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Bot online")

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/api/generate", generateHandler)
	http.HandleFunc("/site/", siteHandler)
	http.HandleFunc("/api/sites", sitesHandler)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	log.Println("Server running on port", port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func onMessage(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.Bot {
		return
	}

	content := strings.TrimSpace(m.Content)

	if content == "" {
		return
	}

	s.ChannelTyping(m.ChannelID)

	go func() {

		site := generateSite(content)

		saveSite(site)

		msg := fmt.Sprintf(
			"✅ Сайт создан:\n%s/site/%s",
			domain,
			site.ID,
		)

		_, err := s.ChannelMessageSend(
			m.ChannelID,
			msg,
		)

		if err != nil {
			log.Println(err)
		}

	}()
}

func generateHandler(w http.ResponseWriter, r *http.Request) {

	type Req struct {
		Prompt string `json:"prompt"`
	}

	var req Req

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	site := generateSite(req.Prompt)

	saveSite(site)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(site)
}

func generateSite(prompt string) Site {

	id := randomID()

	niche := detectNiche(prompt)

	name := generateName(niche)

	description := generateDescription(niche)

	html := buildHTML(name, niche, description)

	return Site{
		ID:          id,
		Name:        name,
		Niche:       niche,
		Description: description,
		HTML:        html,
		URL:         domain + "/site/" + id,
		CreatedAt:   time.Now().Format(time.RFC3339),
	}
}

func detectNiche(prompt string) string {

	p := strings.ToLower(prompt)

	switch {

	case strings.Contains(p, "coffee"):
		return "coffee"

	case strings.Contains(p, "коф"):
		return "coffee"

	case strings.Contains(p, "fit"):
		return "fitness"

	case strings.Contains(p, "фит"):
		return "fitness"

	case strings.Contains(p, "food"):
		return "food"

	case strings.Contains(p, "ед"):
		return "food"

	case strings.Contains(p, "resume"):
		return "resume"

	case strings.Contains(p, "резюме"):
		return "resume"

	default:
		return "business"
	}
}

func generateName(niche string) string {

	data := map[string][]string{
		"coffee": {
			"Coffee House",
			"Dark Roast",
			"Urban Coffee",
		},
		"fitness": {
			"Fit Power",
			"Next Gym",
			"Body Studio",
		},
		"food": {
			"Fast Delivery",
			"Food Point",
			"Food Express",
		},
		"resume": {
			"Resume Pro",
			"CV Builder",
			"Career Boost",
		},
		"business": {
			"Business Pro",
			"Modern Agency",
			"Startup Studio",
		},
	}

	list := data[niche]

	return list[rand.Intn(len(list))]
}

func generateDescription(niche string) string {

	switch niche {

	case "coffee":
		return "Лучший кофе в городе с уютной атмосферой."

	case "fitness":
		return "Современный фитнес клуб для сильных людей."

	case "food":
		return "Быстрая доставка еды за считанные минуты."

	case "resume":
		return "Создавайте профессиональные резюме онлайн."

	default:
		return "Профессиональные решения для бизнеса."
	}
}

func buildHTML(name, niche, description string) string {

	bg := "#0f172a"
	accent := "#7c3aed"

	switch niche {

	case "coffee":
		bg = "#1c1917"
		accent = "#c2410c"

	case "fitness":
		bg = "#111827"
		accent = "#22c55e"

	case "food":
		bg = "#1f2937"
		accent = "#ef4444"

	case "resume":
		bg = "#0f172a"
		accent = "#3b82f6"
	}

	html := fmt.Sprintf(`
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
font-family:Arial,sans-serif;
background:%s;
color:white;
overflow-x:hidden;
}

.hero{
min-height:100vh;
display:flex;
flex-direction:column;
justify-content:center;
align-items:center;
text-align:center;
padding:40px;
}

h1{
font-size:64px;
margin-bottom:20px;
}

p{
font-size:22px;
max-width:700px;
line-height:1.6;
opacity:0.9;
}

button{
margin-top:30px;
padding:16px 32px;
border:none;
border-radius:14px;
background:%s;
color:white;
font-size:18px;
cursor:pointer;
transition:0.3s;
}

button:hover{
transform:scale(1.05);
}

.cards{
display:grid;
grid-template-columns:repeat(auto-fit,minmax(250px,1fr));
gap:20px;
padding:60px;
}

.card{
background:rgba(255,255,255,0.08);
padding:30px;
border-radius:20px;
backdrop-filter:blur(10px);
}

footer{
padding:40px;
text-align:center;
opacity:0.6;
}

</style>

</head>

<body>

<section class="hero">

<h1>%s</h1>

<p>%s</p>

<button>Начать</button>

</section>

<section class="cards">

<div class="card">
<h2>⚡ Быстро</h2>
<p>Современная производительность.</p>
</div>

<div class="card">
<h2>🎨 Красиво</h2>
<p>Профессиональный дизайн.</p>
</div>

<div class="card">
<h2>🚀 Готово</h2>
<p>Мгновенный запуск сайта.</p>
</div>

</section>

<footer>
%s © 2026
</footer>

</body>
</html>
`,
		name,
		bg,
		accent,
		name,
		description,
		name,
	)

	return html
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	html := `
<!DOCTYPE html>
<html>
<head>

<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">

<title>AI Builder</title>

<style>

body{
margin:0;
font-family:Arial,sans-serif;
background:#0f172a;
color:white;
}

.hero{
height:100vh;
display:flex;
flex-direction:column;
justify-content:center;
align-items:center;
text-align:center;
padding:20px;
}

h1{
font-size:72px;
margin-bottom:20px;
}

p{
font-size:22px;
opacity:0.8;
max-width:700px;
}

a{
margin-top:30px;
padding:18px 36px;
background:#7c3aed;
border-radius:16px;
text-decoration:none;
color:white;
font-size:18px;
}

</style>

</head>

<body>

<section class="hero">

<h1>🚀 AI Builder</h1>

<p>Create websites instantly with smart generation.</p>

<a href="/api/sites">
Open Sites
</a>

</section>

</body>
</html>
`

	w.Write([]byte(html))
}

func siteHandler(w http.ResponseWriter, r *http.Request) {

	id := strings.TrimPrefix(r.URL.Path, "/site/")

	sites := loadSites()

	for _, site := range sites {

		if site.ID == id {

			tmpl, err := template.New("site").Parse(site.HTML)

			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}

			var buf bytes.Buffer

			err = tmpl.Execute(&buf, nil)

			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}

			w.Header().Set("Content-Type", "text/html")

			w.Write(buf.Bytes())

			return
		}
	}

	http.NotFound(w, r)
}

func sitesHandler(w http.ResponseWriter, r *http.Request) {

	sites := loadSites()

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(sites)
}

func saveSite(site Site) {

	sites := loadSites()

	sites = append(sites, site)

	data, _ := json.MarshalIndent(sites, "", "  ")

	os.WriteFile(sitesFile, data, 0644)
}

func loadSites() []Site {

	var sites []Site

	file, err := os.Open(sitesFile)

	if err != nil {
		return []Site{}
	}

	defer file.Close()

	data, _ := io.ReadAll(file)

	json.Unmarshal(data, &sites)

	return sites
}

func randomID() string {

	chars := "abcdefghijklmnopqrstuvwxyz0123456789"

	var result strings.Builder

	for i := 0; i < 12; i++ {
		result.WriteByte(chars[rand.Intn(len(chars))])
	}

	return result.String()
}
