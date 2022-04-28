package main

import (
	"fmt"

	sm "github.com/DNA-string-matching/string_matcher"
	ss "github.com/DNA-string-matching/string_similarity"
)

func printResult(method string, index int) {
	if index != -1 {
		fmt.Printf("Matched with %s method in index: %d\n", method, index)
	} else {
		fmt.Printf("Not matched with %s method\n", method)
	}
}

func DNA_String_Matching(pattern, text string, algoIndex int) (float32, bool) {

	similarityDict := make(map[string]float32)
	var index int
	var count int = 0
	var c = make(chan int)

	//pattern := "AGCTGA"
	//text := "AGCTAGCATAAGCTAGCTA"
	//algoIndex := 2

	/* Memilih Algoritma */
	switch algoIndex {
	case 0:
		go sm.BMooreMatcher(pattern, text, c, &index)
	case 1:
		go sm.KMPMatcher(pattern, text, c, &index)
	case 2:
		go sm.BruteForceMatching(pattern, text, c, &index)
	case 3:
		go sm.RegexMatch(pattern, text, c, &index)
	}
	go ss.SmithWatermanSimilarity(pattern, text, c, &similarityDict)

	/* Channel Message Receiver */
	for {
		msg := <-c
		if msg != -1 {
			count++
		}

		if count == 2 {
			// Total 2 Proses
			// Mencari index dan similarity
			break
		}
	}

	/* Check if is Positive */
	isPositive := false
	similarity := ss.FindMaxSimilarity(similarityDict)

	if index != -1 || similarity > 80 {
		isPositive = true
	}

	/* Return */
	return similarity, isPositive

}
