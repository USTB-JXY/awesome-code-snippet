package leetcode

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := make([]int, 0, len(tokens))
	for _, token := range tokens {
		v, err := strconv.Atoi(token)
		if err == nil {
			stack = append(stack, v)
		} else {
			num1, num2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			switch token {
			case "+":
				stack = append(stack, num1+num2)
			case "-":
				stack = append(stack, num1-num2)
			case "*":
				stack = append(stack, num1*num2)
			case "/":
				stack = append(stack, num1/num2)
			}
		}
	}
	return stack[0]
}
func myevalRPN(tokens []string) int {
	stack := []int{}
	for _, token := range tokens {
		v, err := strconv.Atoi(token)
		if err == nil {
			stack = append(stack, v)
		} else {
			n1, n2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			switch token {
			case "+":
				stack = append(stack, n1+n2)
			case "-":
				stack = append(stack, n1-n2)
			case "*":
				stack = append(stack, n1*n2)
			case "/":
				stack = append(stack, n1/n2)
			}
		}
	}
	return stack[0]
}
