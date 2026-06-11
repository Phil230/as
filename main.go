package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("usage: go run . \"text\"")
		return
	}

	text := os.Args[1]

	_, err := ValidateInput(text)
	if err != nil {
		fmt.Println(err)
		return
	}

	banner, err := LoadBanner("standard.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	art := GenerateArt(text, banner)

	fmt.Print(art)
}