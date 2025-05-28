// package main

// import (
// 	"fmt"
// 	"jm-CICO/config"
// 	"jm-CICO/handlers"
// 	"log"
// 	"net/http"
// )

// func main() {
// 	for location, cfg := range config.Locations {
		
// 		// Register endpoint utama
// 		endpoint := "/api/data/" + location
// 		http.HandleFunc(endpoint, handlers.LocationHandlerFactory(cfg))

// 		// Register endpoint subtotal
// 		subtotalEndpoint := endpoint + "/subtotal"
// 		http.HandleFunc(subtotalEndpoint, handlers.LocationHandlerFactory(cfg))
// 	}

// 	for cabang, cfg := range config.CabangDetails {
// 		endpoint := "/api/cabang/" + cabang
// 		http.HandleFunc(endpoint, handlers.DetailHandlerFactory(cfg))
// 	}

// 	fmt.Println("Server running at http://localhost:8080")
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }