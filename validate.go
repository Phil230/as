package main

import (
	"strings"
	"errors"
)

func ValidateInput(text string)(rune, error) {

	input :=  strings.Split(text, "\\n")

	for _, word := range input {
		for _, char := range word {
			if char < 32 || char > 126 {
				return char, errors.New("invalid input")
			}
		}
	}
	return 0, nil
}