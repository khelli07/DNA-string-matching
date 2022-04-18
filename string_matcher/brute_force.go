package string_matcher

func BruteForceMatching(pattern, text string) int {
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
		return i - j
	}

	return -1
}
