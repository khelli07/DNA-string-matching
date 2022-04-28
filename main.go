package main

import (
	"fmt"

	sm "tubes03/string_matcher"
	ss "tubes03/string_similarity"
)

func printResult(method string, index int) {
	if index != -1 {
		fmt.Printf("Matched with %s method in index: %d\n", method, index)
	} else {
		fmt.Printf("Not matched with %s method\n", method)
	}
}

func main() {
	similarityDict1 := make(map[string]float32)
	similarityDict2 := make(map[string]float32)
	var index1 int
	var index2 int
	var c1 = make(chan string)
	var c2 = make(chan string)

	pattern := "TADDC"
	//pattern2 := "TAG"
	text := "AGCTAGCATAAGCTAGCTA"

	ss.SmithWatermanSimilarity(pattern, text)

	go sm.BMooreMatcher(pattern, text, c1, &index1, &similarityDict1)
	go sm.KMPMatcher(pattern, text, c2, &index2, &similarityDict2)

	fmt.Println("Progress")

	// First algo
	for {
		msg := <-c1
		if msg == "Done" {
			fmt.Println("Bmoore")
			if index1 == -1 {
				similarity, text := ss.FindMaxSimilarity(similarityDict1)

				fmt.Println("Closest Text :", text, "with similarity of", similarity, "%")
			}
			if index1 != -1 {
				fmt.Println(index1)
			}
			break
		}
	}

	// Second algo
	for {
		msg := <-c2
		if msg == "Done" {
			fmt.Println("KMP")
			if index2 == -1 {
				similarity, text := ss.FindMaxSimilarity(similarityDict2)

				fmt.Println("Closest Text :", text, "with similarity of", similarity, "%")
			}
			if index2 != -1 {
				fmt.Println(index2)
			}
			break
		}
	}

	//fmt.Printf("%f", ss.LevenshteinSimilarity(pattern, pattern2))

	return

	// RESULT ARE IN INDEX

	printResult("Brute Force", sm.BruteForceMatching(pattern, text))
	printResult("Regex", sm.RegexMatch(pattern, text))
}
