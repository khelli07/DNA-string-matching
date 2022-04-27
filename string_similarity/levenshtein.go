package string_similarity

func findMin(a [3]int) (min int) {
	min = a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
	}
	return min
}

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
	DynamicFilling(&matrix, pattern, text)
	//fmt.Println(matrix)

	similarity = float32(m-1-matrix[m-1][m-1]) / float32(m-1)
	return similarity * 100
}

func DynamicFilling(matrix *[][]int, pattern string, text string) {

	for indexI, row := range (*matrix)[1:] {
		for indexJ, _ := range row[1:] {
			// Find min value
			values := [3]int{
				(*matrix)[indexI][indexJ],
				(*matrix)[indexI][indexJ+1],
				(*matrix)[indexI+1][indexJ],
			}

			temp := isCharDiff(pattern[indexI], text[indexJ])

			//fmt.Println(temp)
			(*matrix)[indexI+1][indexJ+1] = findMin(values) + temp

		}
	}
}

func isCharDiff(pattern, text byte) int {
	if pattern != text {
		return 1
	}

	return 0
}
