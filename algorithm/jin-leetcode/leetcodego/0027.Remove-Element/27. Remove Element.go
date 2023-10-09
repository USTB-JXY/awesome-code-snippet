package leetcode

import "fmt"

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
			}
			j++
		}
	}
	return j
}
func reE(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
				fmt.Print(i, j, " ")
				fmt.Println(nums)
			}
			j++
		}

	}
	fmt.Println(nums)
	return j
}

// 【input】:{[0 1 2 2 3 0 4 2] 2}  【ans】:{5}
// 4 2 [0 1 3 2 2 0 4 2]
// 5 3 [0 1 3 0 2 2 4 2]
// 6 4 [0 1 3 0 4 2 2 2]
// [0 1 3 0 4 2 2 2]
// 【output】:5
// j前面的都是不等于val ，j指向第一个等于val
// 等于val时 j不加
