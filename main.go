package main

import (
	"fmt"
	"jm-CICO/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/loketcabangjawatengah", handlers.LoketCabangHandler)
	http.HandleFunc("/samsatkendal", handlers.SamsatKendalHandler)
	http.HandleFunc("/samsatdemak", handlers.SamsatDemakHandler)

	http.HandleFunc("/samsatpurwodadi", handlers.SamsatPurwodadiHandler)
	http.HandleFunc("/samsatungaran", handlers.SamsatUngaranHandler)
	http.HandleFunc("/samsatsalatiga", handlers.SamsatSalatigaHandler)
	http.HandleFunc("/samsatlokperwsra", handlers.SamsatLokPerwSRAHandler)
	http.HandleFunc("/samsatsurakarta", handlers.SamsatSurakartaHandler)

	fmt.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}