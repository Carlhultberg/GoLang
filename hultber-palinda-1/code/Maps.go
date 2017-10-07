
package main


import (
"golang.org/x/tour/wc"
"strings"
)

func WordCount(s string) map[string]int {
	string := make(map[string]int)
	splittedStrings := strings.Fields(s)
	//use "_" to get subset
	for _, splittedString := range splittedStrings {
		string[splittedString] += 1
	}
	return string
}

func main() {
	wc.Test(WordCount)
}
