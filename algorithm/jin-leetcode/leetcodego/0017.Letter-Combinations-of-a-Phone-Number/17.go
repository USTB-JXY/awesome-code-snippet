package leetcode

import "fmt"

// 解法三 回溯（参考回溯模板，类似DFS）

func myletterCombinationsBT(digits string) []string {
	if digits == "" {
		return []string{}
	}
	result := []string{}
	var dict = map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}
	res := ""
	n := len(digits)
	myletterFunc(&result, dict, res, digits, n)
	fmt.Println(result)
	return result
}

func myletterFunc(result *[]string, dict map[string][]string, res string, digits string, n int) {
	fmt.Println(res)
	if len(res) == n {
		*result = append(*result, res)
		fmt.Println(result)
		return
	}
	k := digits[0:1]
	digits = digits[1:]
	for i := 0; i < len(dict[k]); i++ {
		res += dict[k][i]
		myletterFunc(result, dict, res, digits, n)
		res = res[0 : len(res)-1]
	}
}

func myletterCombinationsBT1(digits string) []string {
	if digits == "" {
		return []string{}
	}
	result := []string{}
	var dict = map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}
	res := ""
	n := len(digits)
	myletterFunc1(result[:], dict, res, digits, n)
	fmt.Println(result)
	return result
}

func myletterFunc1(result []string, dict map[string][]string, res string, digits string, n int) {
	//fmt.Println(res)
	if len(res) == n {
		result = append(result, res)
		//fmt.Println(result)
		fmt.Println(&result)
		return
	}
	k := digits[0:1]
	digits = digits[1:]
	for i := 0; i < len(dict[k]); i++ {
		res += dict[k][i]
		fmt.Println(&result)
		myletterFunc1(result[:], dict, res, digits, n)
		res = res[0 : len(res)-1]
	}
}
