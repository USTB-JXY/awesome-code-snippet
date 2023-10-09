package leetcode

import (
	"fmt"
	"strings"
)

func reverseWords151(s string) string {
	fmt.Println(s)
	ss := strings.Fields(s)
	fmt.Println(ss)
	reverse151(&ss, 0, len(ss)-1)
	return strings.Join(ss, " ")
}

func reverse151(m *[]string, i int, j int) {
	for i <= j {
		(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
		i++
		j--
	}
}

func myreverseWords151(s string) string {
	ss := strings.Fields(s)
	myreverse151(&ss, 0, len(ss)-1)
	return strings.Join(ss, " ")
}

func myreverse151(m *[]string, i int, j int) {
	for i < j {
		(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
		i++
		j--
	}
}
