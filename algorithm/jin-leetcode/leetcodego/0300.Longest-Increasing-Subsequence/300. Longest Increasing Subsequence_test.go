package leetcode

import (
	"fmt"
	"testing"
)

type question300 struct {
	para300
	ans300
}

// para 是参数
// one 代表第一个参数
type para300 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans300 struct {
	one int
}

func Test_Problem300(t *testing.T) {

	qs := []question300{

		{
			para300{[]int{10, 9, 2, 5, 3, 7, 101, 18}},
			ans300{4},
		},
		{
			para300{[]int{1, 2, 3, 5, 99, 10, 3, 4, 5, 6, 7, 8, 9, 10}},
			ans300{10},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 300------------------------\n")

	for _, q := range qs {
		_, p := q.ans300, q.para300
		fmt.Printf("【input】:%v       【output】:%v\n", p, mylengthOfLIS1(p.one))
	}
	fmt.Printf("\n\n\n")
}
