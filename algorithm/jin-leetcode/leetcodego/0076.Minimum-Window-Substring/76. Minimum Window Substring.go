package leetcode

func minWindow(s string, t string) string {
	if s == "" || t == "" {
		return ""
	}
	var tFreq, sFreq [256]int
	result, left, right, finalLeft, finalRight, minW, count := "", 0, -1, -1, -1, len(s)+1, 0

	for i := 0; i < len(t); i++ {
		tFreq[t[i]-'a']++
	}

	for left < len(s) {
		if right+1 < len(s) && count < len(t) {
			sFreq[s[right+1]-'a']++
			if sFreq[s[right+1]-'a'] <= tFreq[s[right+1]-'a'] {
				count++
			}
			right++
		} else {
			if right-left+1 < minW && count == len(t) {
				minW = right - left + 1
				finalLeft = left
				finalRight = right
			}
			if sFreq[s[left]-'a'] == tFreq[s[left]-'a'] {
				count--
			}
			sFreq[s[left]-'a']--
			left++
		}
	}
	if finalLeft != -1 {
		result = string(s[finalLeft : finalRight+1])
	}
	return result
}

func myminWindow(s string, t string) string {
	if s == "" || t == "" {
		return ""
	}
	count := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		count[t[i]]++
	}
	res := s
	for i := 0; i < len(s); i++ {
		size := 0
		tempcount := make(map[byte]int)
		for left, right := i, i; left < len(s) && right < len(s); {
			if size >= len(t) && right-left < len(res) {
				res = s[left:right]
			}
			if count[s[right]] > 0 && tempcount[s[right]] < count[s[right]] {
				tempcount[s[right]]++
				size++
				right++
			} else {
				if count[s[left]] > 0 {
					tempcount[s[left]]--
					size--
				}
				left++
			}
		}

	}
	return res
}
