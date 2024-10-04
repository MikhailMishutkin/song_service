package main

import (
	"github.com/joho/godotenv"
	"log"
	"song_service/configs"
	"song_service/internal/app"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

//	@title			Song Base
//	@version		1.0
//	@description	API to give songs info

//	@host		localhost:9000
//	@BasePath	/

func main() {
	config, err := configs.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	if err := app.StartBase(*config); err != nil {
		log.Fatal(err)
	}
}
