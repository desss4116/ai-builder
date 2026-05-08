package storage

import (
	"encoding/json"
	"os"
)

type Site struct {
	ID    string `json:"id"`
	HTML  string `json:"html"`
	Title string `json:"title"`
}

func SaveSite(site Site) error {

	file, err := os.ReadFile("data/sites.json")
	if err != nil {
		return err
	}

	var sites []Site

	json.Unmarshal(file, &sites)

	sites = append(sites, site)

	updated, _ := json.MarshalIndent(sites, "", "  ")

	return os.WriteFile("data/sites.json", updated, 0644)
}
