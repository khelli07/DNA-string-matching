package string_matcher

import regex "regexp"

func RegexMatch(pattern, text string, c chan int, index *int) {
	re, _ := regex.Compile(pattern)
	idx := re.FindStringIndex(text)

	if len(idx) == 0 {
		*index = -1
	} else {
		*index = idx[0]
	}

	c <- 4
	return
}
