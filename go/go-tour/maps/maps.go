package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	wordCount := make(map[string]int)
	for _, v := range strings.Fields(s) {
		if _, ok := wordCount[v]; ok {
			wordCount[v] += 1
		} else {
			wordCount[v] = 1
		}
	}
	return wordCount
}

func main() {
	wc.Test(WordCount)
}
