package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var longeur int

	for {
		fmt.Print("Entrez la longueur: ")
		_, err := fmt.Scan(&longeur)
		if err == nil {
			break
		}
	}

	chars := []rune("abcdefghijklmnopqrstuvwxyz0123456789,?;.:/!§ù%*µ^¨$£)°=+")
	passw := make([]rune, longeur)

	for i := range passw {
		passw[i] = chars[rand.Intn(len(chars))]
	}

	fmt.Println(string(passw))
}
