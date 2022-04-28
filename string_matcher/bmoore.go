package string_matcher

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

func BMooreMatcher(pattern, text string, c chan int, index *int) {
	n := len(text)
	m := len(pattern)

	if n < m {
		*index = -1
		c <- 1
		return
	}

	mapper := bmMapper(pattern)
	i := m - 1
	j := m - 1

	for i < n {
		if text[i] == pattern[j] {
			if j == 0 {
				*index = i
				c <- 1
				return
			}
			i--
			j--
		} else {
			locc := mapper[text[i]]
			i = i + m - min(j, 1+locc)
			j = m - 1
		}
	}

	*index = -1

	c <- 1

	return
}
