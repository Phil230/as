package main

import "strings"

func RenderLine(text string, fontmap map[rune][]string) []string {
	result := []string{}

	for row := 0; row < 8; row++ {
		var line strings.Builder
		for _, char := range text {
			line.WriteString(fontmap[char][row])
		}
		result = append(result, line.String())
	}
	return result
}