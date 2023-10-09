package leetcode

import "strings"

func wordPattern(pattern string, str string) bool {
	strList := strings.Split(str, " ")
	patternByte := []byte(pattern)
	if pattern == "" || len(patternByte) != len(strList) {
		return false
	}
	pMap := map[byte]string{}
	sMap := map[string]byte{}
	for index, b := range patternByte {
		if _, ok := pMap[b]; !ok {
			if _, ok = sMap[strList[index]]; !ok {
				pMap[b] = strList[index]
				sMap[strList[index]] = b
			} else {
				if sMap[strList[index]] != b {
					return false
				}
			}
		} else {
			if pMap[b] != strList[index] {
				return false
			}
		}
	}
	return true
}
func mywordPattern(pattern string, str string) bool {
	steList := strings.Split(str, " ")
	if len(steList) != len(pattern) {
		return false
	}
	psMap, spMap := map[byte]string{}, map[string]byte{}
	for i := 0; i < len(pattern); i++ {
		_, psok := psMap[pattern[i]]
		_, spok := spMap[steList[i]]
		if !psok && !spok {
			psMap[pattern[i]] = steList[i]
			spMap[steList[i]] = pattern[i]
			continue
		}
		// Input:pattern = "abba", str = "dog cat dog fish"
		// Output: false
		if (psok && spok) && psMap[pattern[i]] == steList[i] {
			continue
		}
		return false
	}
	return true
}
