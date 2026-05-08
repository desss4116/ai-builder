package storage

import (
	"encoding/json"
	"os"
)

type Site struct {
	ID   string `json:"id"`
	HTML string `json:"html"`
}

func SaveSite(site Site) error {

	var sites []Site

	file, err := os.ReadFile(
		"data/sites.json",
	)

	if err == nil {
		json.Unmarshal(file, &sites)
	}

	sites = append(sites, site)

	updated, _ := json.MarshalIndent(
		sites,
		"",
		"  ",
	)

	return os.WriteFile(
		"data/sites.json",
		updated,
		0644,
	)
}
