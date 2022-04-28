package string_similarity

func SmithWatermanSimilarity(pattern, text string, c chan int, similarityDict *map[string]float32) {
	m := len(text) + 1
	n := len(pattern) + 1

	// Initiating Matrix
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, m)
	}

	//fmt.Println(matrix)
	maxValue := DynamicFillingWaterman(&matrix, pattern, text)

	calcSimilarityWaterman(similarityDict, text, pattern, matrix, maxValue)

	c <- 5
}

func DynamicFillingWaterman(matrix *[][]int, pattern string, text string) int {
	maxValue := 0

	for indexI, row := range (*matrix)[1:] {
		for indexJ, _ := range row[1:] {

			values := [3]int{
				max(0, (*matrix)[indexI][indexJ]+isMatch(pattern[indexI], text[indexJ])),
				max(0, (*matrix)[indexI][indexJ+1]-2),
				max(0, (*matrix)[indexI+1][indexJ]-2),
			}

			//fmt.Println(temp)
			(*matrix)[indexI+1][indexJ+1] = findMax(values)

			if (*matrix)[indexI+1][indexJ+1] > maxValue {
				maxValue = (*matrix)[indexI+1][indexJ+1]
			}

		}
	}

	return maxValue
}
