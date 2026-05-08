package storage

import (
	"encoding/json"
	"os"
)

type Memory struct {
	UserID string `json:"user_id"`
	Query  string `json:"query"`
}

func SaveMemory(memory Memory) error {

	file, err := os.ReadFile("data/memory.json")
	if err != nil {
		return err
	}

	var memories []Memory

	json.Unmarshal(file, &memories)

	memories = append(memories, memory)

	updated, _ := json.MarshalIndent(memories, "", "  ")

	return os.WriteFile("data/memory.json", updated, 0644)
}
