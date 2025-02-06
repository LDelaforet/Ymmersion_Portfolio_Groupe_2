package main

import (
    "fmt"
    "math/rand"
)

func main() {
    target := rand.Intn(100) + 1
    var guess int
    for {
        fmt.Print("Entrez votre nombre: ")
        fmt.Scan(&guess)

        if guess < target {
            fmt.Println("C'est plus !")
        } else if guess > target {
            fmt.Println("C'est moins !")
        } else {
            fmt.Println("Bravo")
            break
        }
    }
}