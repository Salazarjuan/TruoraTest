package main

import (
	"Projects/TruoraTest/systems/app"
	DB "Projects/TruoraTest/systems/db"

	"flag"
	"os"

	"github.com/joho/godotenv"
)

var port string

func init() {
	flag.StringVar(&port, "port", "8000", "Assigning the port that the servar should listen on")
	flag.Parse()

	if err := godotenv.Load("config.ini"); err != nil {
		panic(err)
	}

	envPort := os.Getenv("PORT")
	if len(envPort) > 0 {
		port = envPort
	}
}

func main() {
	_, err := DB.Conect()
	if err != nil {
		panic(err)
	}

	s := app.NewServer()

	s.Init(port, db)
	s.Start()
}
