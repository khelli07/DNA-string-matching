package string_matcher

import regex "regexp"

func RegexMatch(pattern, text string) int {
	re, _ := regex.Compile(pattern)
	idx := re.FindStringIndex(text)

	if len(idx) == 0 {
		return -1
	} else {
		return idx[0]
	}
}
