package string_similarity

import "fmt"

func SmithWatermanSimilarity(pattern, text string) float32 {
	similarity := float32(0)
	m := len(text) + 1
	n := len(pattern) + 1

	// Initiating Matrix
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, m)
	}

	//fmt.Println(matrix)
	printMatrix(matrix)
	maxValue := DynamicFillingWaterman(&matrix, pattern, text)
	printMatrix(matrix)
	fmt.Println(maxValue)

	similarityDict := make(map[string]float32)
	calcSimilarityWaterman(&similarityDict, text, pattern, matrix, maxValue)
	fmt.Println(similarityDict)

	similarity = float32(m-1-matrix[m-1][m-1]) / float32(m-1)
	return similarity * 100
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
