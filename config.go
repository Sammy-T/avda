package main

import (
	_ "embed"
	"encoding/json"
	"log"
)

type config struct {
	Info info `json:"info"`
}

type info struct {
	Version string `json:"productVersion"`
}

//go:embed wails.json
var configData []byte

var appInfo info

func init() {
	var cfg config

	err := json.Unmarshal(configData, &cfg)
	if err != nil {
		log.Fatal(err)
	}

	appInfo = cfg.Info
}
