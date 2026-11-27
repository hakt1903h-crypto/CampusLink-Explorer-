package main

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"time"
)

// Les données envoyées aux gabarits (champs en majuscule, comme dans le cours).
type PageData struct {
	Titre string
}

var pages = map[string]*template.Template{}

// Chaque page = le gabarit de base + son propre fichier.
func chargerPages() {
	for _, nom := range []string{"accueil"} {
		pages[nom] = template.Must(template.ParseFiles(
			"templates/base.html",
			"templates/"+nom+".html",
		))
	}
}

func afficher(w http.ResponseWriter, nom string, statut int, data PageData) {
	var buf bytes.Buffer
	if err := pages[nom].ExecuteTemplate(&buf, "base", data); err != nil {
		log.Println("erreur de gabarit :", err)
		http.Error(w, "Erreur interne", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(statut)
	buf.WriteTo(w)
}

func accueil(w http.ResponseWriter, r *http.Request) {
	afficher(w, "accueil", http.StatusOK, PageData{Titre: "Accueil"})
}

func main() {
	chargerPages()

	mux := http.NewServeMux()
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	mux.HandleFunc("GET /{$}", accueil)

	srv := &http.Server{
		Addr:              "localhost:8080",
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Println("Serveur démarré sur http://localhost:8080")
	log.Fatal(srv.ListenAndServe())
}
