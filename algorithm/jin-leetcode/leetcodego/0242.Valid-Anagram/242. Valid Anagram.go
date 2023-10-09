package leetcode

// 解法一
func isAnagram(s string, t string) bool {
	alphabet := make([]int, 26)
	sBytes := []byte(s)
	tBytes := []byte(t)
	if len(sBytes) != len(tBytes) {
		return false
	}
	for i := 0; i < len(sBytes); i++ {
		alphabet[sBytes[i]-'a']++
	}
	for i := 0; i < len(tBytes); i++ {
		alphabet[tBytes[i]-'a']--
	}
	for i := 0; i < 26; i++ {
		if alphabet[i] != 0 {
			return false
		}
	}
	return true
}
func myisAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	alaArr := make([]int, 26)
	for i := 0; i < len(s); i++ {
		alaArr[s[i]-'a']++
	}
	for i := 0; i < len(t); i++ {
		alaArr[t[i]-'a']--
	}
	for i := 0; i < len(alaArr); i++ {
		if alaArr[i] != 0 {
			return false
		}
	}
	return true
}

// 解法二
func isAnagram1(s string, t string) bool {
	hash := map[rune]int{}
	for _, value := range s {
		hash[value]++
	}
	for _, value := range t {
		hash[value]--
	}
	for _, value := range hash {
		if value != 0 {
			return false
		}
	}
	return true
}
