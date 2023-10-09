package leetcode

import "fmt"

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	res := []string{}
	findGenerateParenthesis(n, n, "", &res)
	return res
}

func findGenerateParenthesis(lindex, rindex int, str string, res *[]string) {
	fmt.Println(str)
	// 保证只有n个( n个)
	if lindex == 0 && rindex == 0 {
		*res = append(*res, str)
		return
	}
	if lindex > 0 {
		findGenerateParenthesis(lindex-1, rindex, str+"(", res)
	}
	// if rindex > 0 && lindex < rindex {
	// 	findGenerateParenthesis(lindex, rindex-1, str+")", res)
	// }
	//对每个str 的子集 ( 数量大于等于 )
	if rindex > 0 && lindex < rindex {
		findGenerateParenthesis(lindex, rindex-1, str+")", res)
	}
}

func mygenerateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	res := []string{}
	gcres(n, n, "", &res)
	return res
}

func gcres(n1, n2 int, str string, res *[]string) {
	if n1 == 0 && n2 == 0 {
		*res = append(*res, str)
	}
	if n1 > 0 {
		gcres(n1-1, n2, str+"(", res)
	}
	if n2 > 0 && n2 > n1 {
		gcres(n1, n2-1, str+")", res)
	}
}
