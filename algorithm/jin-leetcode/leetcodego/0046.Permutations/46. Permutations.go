package leetcode

import "fmt"

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	used, p, res := make([]bool, len(nums)), []int{}, [][]int{}
	generatePermutation(nums, 0, p, &res, &used)
	return res
}

func generatePermutation(nums []int, index int, p []int, res *[][]int, used *[]bool) {
	if index == len(nums) {
		temp := make([]int, len(p))
		copy(temp, p)
		*res = append(*res, temp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if !(*used)[i] {
			(*used)[i] = true
			p = append(p, nums[i])
			generatePermutation(nums, index+1, p, res, used)
			p = p[:len(p)-1]
			(*used)[i] = false
		}
	}
	return
}

func mypermute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	c, res, used := []int{}, [][]int{}, make([]bool, len(nums))
	mygenerateCombinations(nums, 0, c, &res, &used)
	return res
}
func mygenerateCombinations(n []int, start int, c []int, res *[][]int, used *[]bool) {
	if len(c) == len(n) {
		b := make([]int, len(c))
		copy(b, c)
		*res = append(*res, b)
		fmt.Println(*res)
		return
	}
	for i := start; i < len(n); i++ {
		if !(*used)[i] {
			c = append(c, n[i])
			(*used)[i] = true
			mygenerateCombinations(n, 0, c, res, used)
			c = c[:len(c)-1]
			(*used)[i] = false
		}

	}
	return
}
