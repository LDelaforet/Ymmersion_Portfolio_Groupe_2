package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type PassedStruct struct {
	State1 string
	State2 string
	State3 string
	State4 string
	State5 string
	State6 string
}

type FlipRequest struct {
	CardID string `json:"cardId"`
}

func main() {
	PassedVars := PassedStruct{}

	// Charger les templates HTML
	temp, errTemp := template.ParseGlob("../assets/WebPages/*.html")
	if errTemp != nil {
		fmt.Printf("Error: %v\n", errTemp)
		return
	}

	// Bouton pour aller a /promo
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "index", PassedVars)
	})

	http.HandleFunc("/flipcard", func(w http.ResponseWriter, r *http.Request) {

	})

	// Fichiers statiques
	http.Handle("/Fonts/", http.StripPrefix("/Fonts/", http.FileServer(http.Dir("../assets/Fonts"))))
	http.Handle("/Pictures/", http.StripPrefix("/Pictures/", http.FileServer(http.Dir("../assets/Pictures"))))
	http.Handle("/Sounds/", http.StripPrefix("/Sounds/", http.FileServer(http.Dir("../assets/Sounds"))))
	http.Handle("/Scripts/", http.StripPrefix("/Scripts/", http.FileServer(http.Dir("../assets/Scripts"))))
	http.Handle("/Styles/", http.StripPrefix("/Styles/", http.FileServer(http.Dir("../assets/Styles"))))
	http.Handle("/WebPages/", http.StripPrefix("/webPages/", http.FileServer(http.Dir("../assets/webPages"))))

	// Lancer le serveur
	http.ListenAndServe("127.0.0.1:8080", nil)
}
