package main

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	ChainId int `json:"chainId"`
}

func main() {
	// Read genesis.json
	data, err := os.ReadFile("genesis.json")
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal into Config struct
	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		log.Fatal(err)
	}

	// Modify chainId
	config.ChainId = 8890

	// Marshal modified config back to JSON
	out, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	// Write out.json
	if err := os.WriteFile("out.json", out, 0644); err != nil {
		log.Fatal(err)
	}
}
