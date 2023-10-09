package leetcode

import "fmt"

func search33(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > nums[low] { // 在数值大的一部分区间里
			if nums[low] <= target && target < nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else if nums[mid] < nums[high] { // 在数值小的一部分区间里
			if nums[mid] < target && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		} else {
			if nums[low] == nums[mid] {
				low++
			}
			if nums[high] == nums[mid] {
				high--
			}
		}
	}
	return -1
}

func mysearch33(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	n := len(nums)
	left, right := 0, len(nums)-1
	for left <= right {
		rote := (left + right) / 2
		if nums[rote] > nums[n-1] {
			left = rote + 1
		} else {
			right = rote - 1
		}
	}
	if target <= nums[n-1] {
		right = len(nums) - 1
	} else {
		left, right = 0, left-1
	}
	fmt.Println(left, right)
	for left <= right {
		rote := (left + right) / 2
		if nums[rote] == target {
			return rote
		}
		if nums[rote] > target {
			right = rote - 1
		} else {
			left = rote + 1
		}
	}
	return -1
}
