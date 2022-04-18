package string_matcher

func kmpMapper(pattern string) map[int]int {
	m := make(map[int]int)
	m[0] = 0

	n := len(pattern)
	j := 0
	i := 1
	for i < n {
		if pattern[i] == pattern[j] {
			m[i] = j + 1
			i++
			j++
		} else if j > 0 {
			j = m[j-1]
		} else {
			m[i] = 0
			i++
		}
	}

	return m
}

func KMPMatcher(pattern, text string) int {
	m := len(pattern)
	n := len(text)

	if n < m {
		return -1
	}

	j := 0
	i := 0
	idx := -1

	mapper := kmpMapper(pattern)
	for i < n && j < m {
		if text[i] == pattern[j] {
			if j == m-1 {
				idx = i - j
			}
			i++
			j++
		} else if j > 0 {
			j = mapper[j-1]
		} else {
			i++
		}
	}

	return idx
}
