package main

import "fmt"

func main() {
	// Test the isPalindrome function
	words := []string{"radar", "kayak"}
	for _, word := range words {
		fmt.Printf("%s is palindrome: %v\n", word, isPalindrome(word))
	}
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
