package leetcode

import "fmt"

// 解法一 时间复杂度 O(n)，空间复杂度 O(1)
func rotate(nums []int, k int) {
	k %= len(nums)
	reverse(nums)
	fmt.Println(nums)
	reverse(nums[:k])
	fmt.Println(nums)
	reverse(nums[k:])
	fmt.Println(nums)
}

func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

// 解法二 时间复杂度 O(n)，空间复杂度 O(n)
func rotate1(nums []int, k int) {
	newNums := make([]int, len(nums))
	for i, v := range nums {
		newNums[(i+k)%len(nums)] = v
	}
	copy(nums, newNums)
}

func res(nums []int, k int) {
	k %= len(nums)
	resarray(nums)

	resarray(nums[:k])
	resarray(nums[k:])
}
func resarray(nums []int) {
	i, j := 0, len(nums)-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
