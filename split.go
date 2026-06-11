package main

import "strings"

func SplitInput(text string) []string {
	input := strings.Split(text, "\\n")
	return input
}