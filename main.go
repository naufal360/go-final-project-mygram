package main

import (
	"go-final-project-mygram/app"
	"go-final-project-mygram/config"
	"log"

	"github.com/joho/godotenv"
)

func init() {
	// debug env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error .env file at %s", err)
		panic(err)
	}

	err = config.InitGorm()
	if err != nil {
		panic(err)
	}
}

func main() {
	app.StartApp()
}
