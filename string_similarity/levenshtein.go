package string_similarity

func LevenshteinSimilarity(pattern, text string) float32 {
	similarity := float32(0)
	m := len(pattern) + 1

	// Initiating Matrix
	matrix := make([][]int, m)
	for i := range matrix {
		matrix[i] = make([]int, m)
		matrix[i][0] = i
	}
	for index := range matrix[0] {
		matrix[0][index] = index
	}

	//fmt.Println(matrix)
	DynamicFillingLevenshtein(&matrix, pattern, text)

	similarity = float32(m-1-matrix[m-1][m-1]) / float32(m-1)
	return similarity * 100
}

func DynamicFillingLevenshtein(matrix *[][]int, pattern string, text string) {

	for indexI, row := range (*matrix)[1:] {
		for indexJ, _ := range row[1:] {
			// Find min value
			temp := isCharDiff(pattern[indexI], text[indexJ])

			values := [3]int{
				(*matrix)[indexI][indexJ] + temp,
				(*matrix)[indexI][indexJ+1] + 1,
				(*matrix)[indexI+1][indexJ] + 1,
			}

			//fmt.Println(temp)
			(*matrix)[indexI+1][indexJ+1] = findMin(values)

		}
	}
}
