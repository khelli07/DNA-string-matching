package main

import (
	"fmt"

	sm "github.com/DNA-string-matching/string_matcher"
)

func printResult(method string, index int) {
	if index != -1 {
		fmt.Printf("Matched with %s method in index: %d\n", method, index)
	} else {
		fmt.Printf("Not matched with %s method\n", method)
	}
}

func main() {
	pattern := "AGG"
	text := "AGCTAGCATGCATCGAGG"

	// RESULT ARE IN INDEX
	printResult("Brute Force", sm.BruteForceMatching(pattern, text))
	printResult("KMP", sm.KMPMatcher(pattern, text))
	printResult("Regex", sm.KMPMatcher(pattern, text))
	printResult("Boyer-Moore", sm.BMooreMatcher(pattern, text))

}
