package string_matcher

import (
	"fmt"

	ss "tubes03/string_similarity"
)

func bmMapper(pattern string) map[byte]int {
	m := make(map[byte]int)
	visited := make(map[byte]bool)

	i := len(pattern) - 1
	m['A'] = -1
	m['G'] = -1
	m['C'] = -1
	m['T'] = -1

	for i >= 0 && len(visited) < 4 {
		m[pattern[i]] = i
		visited[pattern[i]] = true
		i--
	}

	return m
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func BMooreMatcher(pattern, text string, c chan string, index *int, similarityDict *map[string]float32) {
	n := len(text)
	m := len(pattern)
	if n < m {
		*index = -1
		return
	}

	mapper := bmMapper(pattern)

	i := m - 1
	j := m - 1

	for i < n {
		if text[i] == pattern[j] {

			if j == 0 {

				fmt.Println(text[i : i+m])
				c <- "Done"
				*index = i
				return
			}
			i--
			j--
		} else {
			currText := text[max(i-m+1, 0) : i+1]
			fmt.Println(1, currText)

			if len(currText) == m {
				(*similarityDict)[currText] = ss.LevenshteinSimilarity(pattern, currText)
			}

			locc := mapper[text[i]]
			i = i + m - min(j, 1+locc)
			j = m - 1
		}
	}

	*index = -1
	c <- "Done"
	return
}
