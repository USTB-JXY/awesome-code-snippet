package leetcode

import "fmt"

func longestCommonPrefix(strs []string) string {
	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		for j := 0; j < len(prefix); j++ {
			if len(strs[i]) <= j || strs[i][j] != prefix[j] {
				prefix = prefix[0:j]
				break
			}
		}
	}

	return prefix
}

func mylongestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for j := 0; j < len(prefix); j++ {
			//flower flow等于j 也应该结束不然下次越界
			if len(strs[i]) <= j || strs[i][j] != prefix[j] {
				prefix = prefix[:j]
				fmt.Println("break")
				break
			}
		}
		fmt.Println("continue next")
	}
	return prefix
}
