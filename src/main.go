package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"
)

type PassedStruct struct {
	Code1 string
	Code2 string
	Code3 string
	Code4 string
	Code5 string
	Code6 string
}

type FlipRequest struct {
	CardID string `json:"cardId"`
}

func ReadEpreuve(epreuve string) string {
	fileContent, err := ioutil.ReadFile("../assets/Epreuves/" + epreuve)
	if err != nil {
		return "Impossible de lire le fichier"
	} else {
		return string(fileContent)
	}
}

func main() {
	CodeEp1 := ""
	CodeEp2 := ""
	CodeEp3 := ""
	CodeEp4 := ""
	CodeEp5 := ""
	CodeEp6 := ""

	// Charger les templates HTML
	temp, errTemp := template.ParseGlob("../assets/WebPages/*.html")
	if errTemp != nil {
		fmt.Printf("Error: %v\n", errTemp)
		return
	}

	CodeEp1 = ReadEpreuve("Epreuve1.css")
	CodeEp2 = ReadEpreuve("Epreuve2.go")
	CodeEp3 = ReadEpreuve("Epreuve3.go")
	CodeEp4 = ReadEpreuve("Epreuve4.py")
	CodeEp5 = ReadEpreuve("Epreuve5.go")
	CodeEp6 = ReadEpreuve("Epreuve6.go")

	PassedVars := PassedStruct{
		Code1: CodeEp1,
		Code2: CodeEp2,
		Code3: CodeEp3,
		Code4: CodeEp4,
		Code5: CodeEp5,
		Code6: CodeEp6,
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
