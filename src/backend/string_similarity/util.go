package string_similarity

func findMax(a [3]int) (max int) {
	max = a[0]
	for _, value := range a {
		if value > max {
			max = value
		}
	}
	return max
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func FindMaxSimilarity(dict map[string]float32) (similarity float32) {
	var maxValue float32 = 0

	for _, value := range dict {
		if value > float32(maxValue) {
			maxValue = value
		}
	}

	return maxValue
}

func calcSimilarityWaterman(dict *map[string]float32, text string, pattern string, matrix [][]int, maxValue int) {
	m := len(pattern)

	for indexi, elements := range matrix {
		for indexj, element := range elements {
			if element == maxValue {
				i := indexi
				j := indexj
				count := 0

				for {
					if matrix[i][j] == 0 {
						break
					}
					count += max(0, matrix[i][j]-matrix[i-1][j-1])

					i--
					j--
				}

				currText := text[j:indexj]
				(*dict)[currText] = float32(count * 100 / m)
			}

		}
	}
}

func isMatch(pattern, text byte) int {
	if pattern != text {
		return 0
	}

	return 1
}
