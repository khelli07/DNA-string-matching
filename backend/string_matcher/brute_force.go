package string_matcher

func BruteForceMatching(pattern, text string, c chan int, index *int) {
	i, j := 0, 0
	for i < len(text) && j < len(pattern) {
		if pattern[j] != text[i] {
			i++
			j = 0
		} else {
			i++
			j++
		}
	}

	if j == len(pattern) {
		*index = i - j
		c <- 2
		return
	}

	*index = -1

	c <- 2
	return
}
