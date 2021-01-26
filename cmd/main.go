package main

import (
	"log"
	server "main/internal/server"
	"os"
)

func init() {
	if err := os.Setenv("ConfigPath", "./config/"); err != nil {
		log.Fatal(err)
	}
}

func main() {
	configPath := os.Getenv("ConfigPath")
	if configPath == "" {
		log.Fatal("env value not set")
	}

	config, err := server.LoadConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}

	db, err := server.ConnectToDB(config.DataBaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	server := server.NewServer(config, db)

	server.Start()
}
