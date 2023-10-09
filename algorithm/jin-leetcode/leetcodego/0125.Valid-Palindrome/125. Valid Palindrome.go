package leetcode

import (
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	i, j := 0, len(s)-1
	for i < j {
		for i < j && !isChar(s[i]) {
			i++
		}
		for i < j && !isChar(s[j]) {
			j--
		}
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

// 判断 c 是否是字符或者数字
func isChar(c byte) bool {
	if ('a' <= c && c <= 'z') || ('0' <= c && c <= '9') {
		return true
	}
	return false
}

func IsChar(c byte) bool {
	if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
		return true
	}
	return false
}

func myisPalindrome(s string) bool {
	s = strings.ToLower(s)
	l, r := 0, len(s)-1
	for l < r {
		for l < r && !IsChar(s[l]) {
			l++
		}
		for l < r && !IsChar(s[r]) {
			r--
		}
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
