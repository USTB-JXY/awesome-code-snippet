package leetcode

import (
	"strconv"
)

func summaryRanges(nums []int) (ans []string) {
	for i, n := 0, len(nums); i < n; {
		left := i
		for i++; i < n && nums[i-1]+1 == nums[i]; i++ {
		}
		s := strconv.Itoa(nums[left])
		if left != i-1 {
			s += "->" + strconv.Itoa(nums[i-1])
		}
		ans = append(ans, s)
	}
	return
}
func mysummaryRanges(nums []int) []string {
	res := []string{}
	for i := 0; i < len(nums); {
		j := i + 1
		for j < len(nums) {
			if nums[j] == nums[j-1]+1 {
				j++
			} else {
				break
			}
		}
		s := strconv.Itoa(nums[i])
		if j != i+1 {
			s += "->" + strconv.Itoa(nums[j-1])
		}
		i = j
		res = append(res, s)
	}
	return res
}
