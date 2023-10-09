package leetcode

func findSubstring(s string, words []string) []int {
	if len(words) == 0 {
		return []int{}
	}
	res := []int{}
	counter := map[string]int{}
	for _, w := range words {
		counter[w]++
	}
	length, totalLen, tmpCounter := len(words[0]), len(words[0])*len(words), resetMap(counter)
	for i, start := 0, 0; i < len(s)-length+1; {
		str := s[i : i+length]
		if tmpCounter[str] > 0 {
			tmpCounter[str]--
			i = i + length
			if i-start == totalLen {
				res = append(res, start)
				//【input】:{aaaaaaaa [aa aa aa]}   【ans】:{[0 1 2]}     【output】:[0 1 2]
				// start++
				// i = start
			}
		} else {
			start++
			i = start
			tmpCounter = resetMap(counter)
		}
	}
	return res
}

func checkWords(s map[string]int) bool {
	flag := true
	for _, v := range s {
		if v > 0 {
			flag = false
			break
		}
	}
	return flag
}

func resetMap(s map[string]int) map[string]int {
	c := map[string]int{}
	for k, v := range s {
		c[k] = v
	}
	return c
}

func myfindSubstring(s string, words []string) []int {
	if len(words) == 0 {
		return []int{}
	}
	res, counter := []int{}, map[string]int{}
	for _, v := range words {
		counter[v]++
	}
	return res
}

func findSubstring1(s string, words []string) []int {
	var result []int
	wordLen := len(words[0])
	wordCount := len(words)
	if wordCount < 1 || len(s) < wordLen*wordCount {
		return result
	}
	dic := make(map[string]int)
	for _, word := range words {
		dic[word]++
	}
	for i := 0; i < wordLen; i++ {
		count := 0
		var tempDic = make(map[string]int)
		for left, right := i, i; right <= len(s)-wordLen; right += wordLen {
			word := s[right : right+wordLen]
			if num, ok := dic[word]; ok {
				for tempDic[word] >= num {
					tempDic[s[left:left+wordLen]]--
					count--
					left = left + wordLen
				}
				tempDic[word]++
				count++
			} else {
				left = right + wordLen
				count = 0
				tempDic = make(map[string]int)
			}
			if count == wordCount {
				result = append(result, left)
			}
		}
	}
	return result
}

func findSubstring2(s string, words []string) (ans []int) {
	ls, m, n := len(s), len(words), len(words[0])
	for i := 0; i < n && i+m*n <= ls; i++ {
		differ := map[string]int{}
		for j := 0; j < m; j++ {
			differ[s[i+j*n:i+(j+1)*n]]++
		}
		for _, word := range words {
			differ[word]--
			if differ[word] == 0 {
				delete(differ, word)
			}
		}
		for start := i; start < ls-m*n+1; start += n {
			if start != i {
				word := s[start+(m-1)*n : start+m*n]
				differ[word]++
				if differ[word] == 0 {
					delete(differ, word)
				}
				word = s[start-n : start]
				differ[word]--
				if differ[word] == 0 {
					delete(differ, word)
				}
			}
			if len(differ) == 0 {
				ans = append(ans, start)
			}
		}
	}
	return
}
