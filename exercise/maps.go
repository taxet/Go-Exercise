package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	var res = make(map[string]int)
	var d = strings.Fields(s)
	for _, v := range d {
		res[v] += 1
	}
	return res
}

func main() {
	wc.Test(WordCount)
}
