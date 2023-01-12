package main

import (
	"fmt"
	"log"
	"moonlay_technical_test/routes"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error read env file with err: %s", err)
	}
	// defer db.Close()
	e := routes.Init()
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", os.Getenv("APP_PORT"))))
}
