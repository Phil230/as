package main

import "strings"

func GenerateArt(text string, bannerMap map[rune][]string) string {
    if text == "" {
        return ""
    }

    text = strings.ReplaceAll(text, "\n", "\\n")
    slice := SplitInput(text)

    allEmpty := true
    for _, w := range slice {
        if w != "" {
            allEmpty = false
            break
        }
    }
    if allEmpty && len(slice) > 0 {
        slice = slice[:len(slice)-1]
    }

    var line strings.Builder

    for _, word := range slice {
        if word == "" {
            line.WriteString("\n")
            continue
        }
        renderLine := RenderLine(word, bannerMap)
        for i := 0; i < 8; i++ {
            line.WriteString(renderLine[i])
			line.WriteByte('\n')
        }
    }

    return line.String()
}