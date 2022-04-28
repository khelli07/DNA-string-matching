package controller

import (
	regex "regexp"
	"strings"

	sm "github.com/DNA-string-matching/backend/string_matcher"
	ss "github.com/DNA-string-matching/backend/string_similarity"
)

func PatternIsValid(pattern string) bool {
	re, _ := regex.Compile("[^AGCT]")
	idx := re.FindStringIndex(pattern)

	if len(idx) != 0 {
		return false
	}

	return true
}

func FindPattern(pattern, text string) bool {
	re, _ := regex.Compile(pattern)
	idx := re.FindStringIndex(text)

	if len(idx) == 0 {
		return false
	} else {
		return true
	}
}

func processDate(date string) string {
	mapper := make(map[string]string)
	mapper["januari"] = "01"
	mapper["februari"] = "02"
	mapper["maret"] = "03"
	mapper["april"] = "04"
	mapper["mei"] = "05"
	mapper["juni"] = "06"
	mapper["juli"] = "07"
	mapper["agustus"] = "08"
	mapper["september"] = "09"
	mapper["oktober"] = "10"
	mapper["november"] = "11"
	mapper["desember"] = "12"

	mapper["january"] = "01"
	mapper["february"] = "02"
	mapper["march"] = "03"
	mapper["may"] = "05"
	mapper["june"] = "06"
	mapper["july"] = "07"
	mapper["august"] = "08"
	mapper["october"] = "10"

	splitted := strings.Split(date, " ")
	splitted[0], splitted[2] = splitted[2], splitted[0]
	splitted[1] = mapper[strings.ToLower(splitted[1])]

	return strings.Join(splitted, "-")
}

func ExtractQuery(query string) (string, string) {
	var ret [2]string

	re1, _ := regex.Compile("(\\d{4}\\-\\d{2}\\-\\d{2})")
	date := re1.FindAllString(query, -1)
	re2, _ := regex.Compile("(\\d{2}.*\\d{4})")

	if len(date) == 0 {
		date = re2.FindAllString(query, -1)
		if len(date) != 0 {
			date[0] = processDate(date[0])
		}
	} else {
		ret[0] = date[0]
	}

	name := re1.ReplaceAllString(query, "")
	name = re2.ReplaceAllString(name, "")

	if strings.TrimSpace(name) != "" {
		ret[1] = strings.TrimSpace(name)
	}

	return ret[0], ret[1]
}

func DNAStringMatching(pattern, text string, algoIndex int) (float32, bool) {

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
