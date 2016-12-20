package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	wordCount := make(map[string]int)
	for _, v := range strings.Fields(s) {
		_, ok := wordCount[v]
		if ok == true {
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
