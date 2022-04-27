package main

import (
	"fmt"

	sm "tubes03/string_matcher"
)

func printResult(method string, index int) {
	if index != -1 {
		fmt.Printf("Matched with %s method in index: %d\n", method, index)
	} else {
		fmt.Printf("Not matched with %s method\n", method)
	}
}

func main() {
	similarityDict := make(map[string]float32)
	var index int
	var c = make(chan string)

	pattern := "AGTT"
	//pattern2 := "TAG"
	text := "AGCTAGCATGCAGGTCGAGG"

	go sm.BMooreMatcher(pattern, text, c, &index, &similarityDict)
	fmt.Println("Progress")

	for {
		msg := <-c
		if msg == "Done" {
			break
		}
	}

	fmt.Println(index)
	fmt.Println(similarityDict)

	//fmt.Printf("%f", ss.LevenshteinSimilarity(pattern, pattern2))

	return

	// RESULT ARE IN INDEX

	printResult("Brute Force", sm.BruteForceMatching(pattern, text))
	printResult("KMP", sm.KMPMatcher(pattern, text))
	printResult("Regex", sm.KMPMatcher(pattern, text))
}
