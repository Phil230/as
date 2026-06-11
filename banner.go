package main

import (
	"errors"
	"os"
	"strings"
)

func LoadBanner(file string) (map[rune][]string, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, errors.New("error reading")
	}
	if len(data) == 0 {
		return nil, errors.New("empty file")
	}

	line := strings.Split(string(data), "\n")

	if len(line) < 855 {
		return nil, errors.New("file corrupt")
	}

	bannerMap := make(map[rune][]string)

	for i := rune(32); i <= 126; i++ {
		start := (int(i) - 32) * 9
		bannerMap[i] = line[start+1 : start+9]
	}
	return bannerMap, nil
}