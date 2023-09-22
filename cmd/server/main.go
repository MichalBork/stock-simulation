package main

import (
	"log"
	"net/http"

	"stock-simulation/api"
)

func main() {
	// Ustaw nasze trasygo run cmd/server/main.go
	api.Routes()

	// Uruchom serwer
	log.Println("Uruchamiam serwer na porcie 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
