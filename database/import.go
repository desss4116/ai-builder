package database

func SavePattern(tag string, content string, source string) {
	stmt, _ := DB.Prepare("INSERT INTO knowledge (tag, content, source) VALUES (?, ?, ?)")
	defer stmt.Close()
	stmt.Exec(tag, content, source)
}

func SaveGeneratedSite(id, prompt, html string) {
	stmt, _ := DB.Prepare("INSERT INTO sites (id, prompt, html) VALUES (?, ?, ?)")
	defer stmt.Close()
	stmt.Exec(id, prompt, html)
}
