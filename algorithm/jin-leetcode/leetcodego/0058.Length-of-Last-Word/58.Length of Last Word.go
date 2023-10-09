package leetcode

func lengthOfLastWord(s string) int {
	last := len(s) - 1
	for last >= 0 && s[last] == ' ' {
		last--
	}
	if last < 0 {
		return 0
	}
	first := last
	for first >= 0 && s[first] != ' ' {
		first--
	}
	return last - first
}

func mylengthOfLastWord(s string) int {
	last := len(s) - 1
	for last >= 0 {
		if s[last] != ' ' {
			break
		}
		last--
	}
	fist := last - 1
	for fist >= 0 {
		if s[fist] == ' ' {
			break
		}
		fist--
	}
	return last - fist
}
