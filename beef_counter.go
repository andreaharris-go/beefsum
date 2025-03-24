package main

import (
	"regexp"
	"strings"
)

var knownMeats = map[string]bool{
	"doner":      true,
	"flank":      true,
	"kielbasa":   true,
	"shoulder":   true,
	"shankle":    true,
	"brisket":    true,
	"pig":        true,
	"biltong":    true,
	"jowl":       true,
	"tongue":     true,
	"ribeye":     true,
	"porchetta":  true,
	"ball":       true,
	"tip":        true,
	"picanha":    true,
	"short":      true,
	"ribs":       true,
	"pancetta":   true,
	"chicken":    true,
	"swine":      true,
	"alcatra":    true,
	"corned":     true,
	"beef":       true,
	"prosciutto": true,
	"ham":        true,
	"hock":       true,
	"loin":       true,
	"pork":       true,
	"boudin":     true,
}

func CountMeats(text string) map[string]int {
	re := regexp.MustCompile(`[^\w\-]+`)
	words := re.Split(strings.ToLower(text), -1)
	counts := make(map[string]int)

	for _, word := range words {
		if knownMeats[word] {
			counts[word]++
		}
	}

	return counts
}
