package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/febriliankr/go-mysql-starter/routes"
	"github.com/rs/cors"
)

func main() {

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	})

	routes.InitializeRoutes()

	hostname := "localhost"
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Listening at http://%s:%s\n", hostname, port)
	err := http.ListenAndServe(":"+port, c.Handler(http.DefaultServeMux))
	if err != nil {
		log.Fatal(err)
	}

}
