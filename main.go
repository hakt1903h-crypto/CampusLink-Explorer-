package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func accueil(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "CampusLink Explorer : Elden Ring")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", accueil)

	srv := &http.Server{
		Addr:              "localhost:8080",
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Println("Serveur démarré sur http://localhost:8080")
	log.Fatal(srv.ListenAndServe())
}
