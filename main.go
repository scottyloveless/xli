package main

import (
	"fmt"
	"os"
)

type apiConfig struct {
	key    string
	secret string
	bearer string
}

func main() {
	// apiKey := os.Getenv("KEY")
	// secret := os.Getenv("SECRET")
	// bearer := os.Getenv("BEARER")
	if len(os.Args) < 2 {
		fmt.Println("Usage: xli \"your post here\"")
		os.Exit(1)
	}

	input := os.Args[1]
	// if input == "login" {
	// 	// LoginHandler()
	// }

	// RefreshToken if needed
	// SendPost

	fmt.Printf("You passed: %s\n", input)
}
