package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Vérification de mot de passe :")
	fmt.Println("Entrez un mdp de 8 caractères ou plus (spériaux inclus):")
	fmt.Println("Entrez exit pour quitter le logiciel")

	for {
		fmt.Print("> ")
		scanner.Scan()
		mdp := scanner.Text()

		if strings.ToLower(mdp) == "exit" {
			fmt.Println("Logiciel fermé")
			break
		}

		strength := checkPasswordStrength(mdp)
		fmt.Printf("Niveau de sécurité : %s\n\n", strength)
	}
}

func checkPasswordStrength(mdp string) string {
	char := "abcdefghijklmnopqrstuvxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	num := "1234567890"
	spe := "!@#$%^&*()-_=+[]{}<>?/"

	hasChar := false
	hasNum := false
	hasSpe := false

	for _, c := range mdp {
		switch {
		case strings.ContainsRune(char, c):
			hasChar = true
		case strings.ContainsRune(num, c):
			hasNum = true
		case strings.ContainsRune(spe, c):
			hasSpe = true
		}
	}

	lenght := len(mdp)
	switch {
	case lenght >= 8 && hasChar && hasNum && hasSpe:
		return "Le mot de passe est fort"
	case lenght >= 8 && hasChar && (hasNum || hasSpe):
		return "Le mot de passe est moyen"
	case lenght >= 8 && hasChar:
		return "Le mot de passe est faible"
	default:
		fmt.Println("Le mot de passe n'est pas assez long ou contiens des caractères invalide")
		return "Pas assez de caractères"
	}
}
