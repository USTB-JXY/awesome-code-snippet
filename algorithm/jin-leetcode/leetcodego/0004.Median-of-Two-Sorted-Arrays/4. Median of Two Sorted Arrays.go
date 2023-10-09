package leetcode

import (
	"fmt"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 假设 nums1 的长度小
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}
	low, high, k, nums1Mid, nums2Mid := 0, len(nums1), (len(nums1)+len(nums2)+1)>>1, 0, 0
	for low <= high {
		// nums1:  ……………… nums1[nums1Mid-1] | nums1[nums1Mid] ……………………
		// nums2:  ……………… nums2[nums2Mid-1] | nums2[nums2Mid] ……………………
		nums1Mid = low + (high-low)>>1 // 分界限右侧是 mid，分界线左侧是 mid - 1
		nums2Mid = k - nums1Mid
		if nums1Mid > 0 && nums1[nums1Mid-1] > nums2[nums2Mid] { // nums1 中的分界线划多了，要向左边移动
			high = nums1Mid - 1
		} else if nums1Mid != len(nums1) && nums1[nums1Mid] < nums2[nums2Mid-1] { // nums1 中的分界线划少了，要向右边移动
			low = nums1Mid + 1
		} else {
			// 找到合适的划分了，需要输出最终结果了
			// 分为奇数偶数 2 种情况
			break
		}
	}
	midLeft, midRight := 0, 0
	if nums1Mid == 0 {
		midLeft = nums2[nums2Mid-1]
	} else if nums2Mid == 0 {
		midLeft = nums1[nums1Mid-1]
	} else {
		midLeft = max(nums1[nums1Mid-1], nums2[nums2Mid-1])
	}
	if (len(nums1)+len(nums2))&1 == 1 {
		return float64(midLeft)
	}
	if nums1Mid == len(nums1) {
		midRight = nums2[nums2Mid]
	} else if nums2Mid == len(nums2) {
		midRight = nums1[nums1Mid]
	} else {
		midRight = min(nums1[nums1Mid], nums2[nums2Mid])
	}
	return float64(midLeft+midRight) / 2
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
func myfindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 假设 nums1 的长度小
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}
	low, high, k, nums1Mid, nums2Mid := 0, len(nums1), (len(nums1)+len(nums2)+1)>>1, 0, 0
	for low <= high {
		// nums1:  ……………… nums1[nums1Mid-1] | nums1[nums1Mid] ……………………
		// nums2:  ……………… nums2[nums2Mid-1] | nums2[nums2Mid] ……………………
		nums1Mid = low + (high-low)>>1 // 分界限右侧是 mid，分界线左侧是 mid - 1
		nums2Mid = k - nums1Mid
		if nums1Mid > 0 && nums1[nums1Mid-1] > nums2[nums2Mid] { // nums1 中的分界线划多了，要向左边移动
			high = nums1Mid - 1
		} else if nums1Mid != len(nums1) && nums1[nums1Mid] < nums2[nums2Mid-1] { // nums1 中的分界线划少了，要向右边移动
			low = nums1Mid + 1
		} else {
			// 找到合适的划分了，需要输出最终结果了
			// 分为奇数偶数 2 种情况
			break
		}
	}
	r1, r2, r3, r4 := nums1Mid+1, nums2Mid+1, nums1Mid, nums2Mid
	if r2 > len(nums2)-1 {
		r2 = len(nums2) - 1
	}
	if r1 > len(nums1)-1 {
		r1 = len(nums1) - 1
	}
	if r4 > len(nums2)-1 {
		r4--
	}
	if r3 > len(nums1)-1 {
		r3--
	}
	minright := min(nums1[r1], nums2[r2])
	maxleft := max(nums1[r3], nums2[r4])
	fmt.Println(maxleft, minright)
	if len(nums1)+len(nums2)%2 == 1 {
		return float64(max(nums1[nums1Mid], nums2[nums2Mid]))
	} else {
		return float64((maxleft + minright) / 2)
	}

}

func gfindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	if totalLength%2 == 1 {
		midIndex := totalLength / 2
		return float64(getKthElement(nums1, nums2, midIndex+1))
	} else {
		midIndex1, midIndex2 := totalLength/2-1, totalLength/2
		return float64(getKthElement(nums1, nums2, midIndex1+1)+getKthElement(nums1, nums2, midIndex2+1)) / 2.0
	}
	return 0
}
func getKthElement(nums1, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for {
		if index1 == len(nums1) {
			return nums2[index2+k-1]
		}
		if index2 == len(nums2) {
			return nums1[index1+k-1]
		}
		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}
		half := k / 2
		newIndex1 := min(index1+half, len(nums1)) - 1
		newIndex2 := min(index2+half, len(nums2)) - 1
		pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
		if pivot1 <= pivot2 {
			k -= (newIndex1 - index1 + 1)
			index1 = newIndex1 + 1
		} else {
			k -= (newIndex2 - index2 + 1)
			index2 = newIndex2 + 1
		}
	}
	return 0
}

func mygfindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totallength := len(nums1) + len(nums2)
	if totallength%2 == 1 {
		return float64(mygetKthElement(nums1, nums2, (totallength+1)/2-1)) / 1.0
	} else {
		return float64(mygetKthElement(nums1, nums2, (totallength)/2)+mygetKthElement(nums1, nums2, (totallength)/2-1)) / 2.0
	}
	return 0
}
func mygetKthElement(nums1, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for {
		if index1 == len(nums1) {
			return nums2[index2+k]
		}
		if index2 == len(nums2) {
			return nums1[index1+k]
		}
		if k == 0 {
			return min(nums1[index1], nums2[index2])
		}
		half := (k + 1) / 2
		nowIndex1 := min(index1+half, len(nums1)) - 1
		nowIndex2 := min(index2+half, len(nums2)) - 1
		pivot1, pivot2 := nums1[nowIndex1], nums2[nowIndex2]
		if pivot1 <= pivot2 {
			k -= (nowIndex1 - index1 + 1)
			index1 = nowIndex1 + 1
		} else {
			k -= (nowIndex2 - index2 + 1)
			index2 = nowIndex2 + 1
		}
	}
	return 0
}

// func min(x, y int) int {
// 	if x < y {
// 		return x
// 	}
// 	return y
// }

// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
// 	if len(nums1) > len(nums2) {
// 		return findMedianSortedArrays(nums2, nums1)
// 	}
// 	m, n := len(nums1), len(nums2)
// 	left, right := 0, m
// 	median1, median2 := 0, 0
// 	for left <= right {
// 		i := (left + right) / 2
// 		j := (m+n+1)/2 - i
// 		nums_im1 := math.MinInt32
// 		if i != 0 {
// 			nums_im1 = nums1[i-1]
// 		}
// 		nums_i := math.MaxInt32
// 		if i != m {
// 			nums_i = nums1[i]
// 		}
// 		nums_jm1 := math.MinInt32
// 		if j != 0 {
// 			nums_jm1 = nums2[j-1]
// 		}
// 		nums_j := math.MaxInt32
// 		if j != n {
// 			nums_j = nums2[j]
// 		}
// 		if nums_im1 <= nums_j {
// 			median1 = max(nums_im1, nums_jm1)
// 			median2 = min(nums_i, nums_j)
// 			left = i + 1
// 		} else {
// 			right = i - 1
// 		}
// 	}
// 	if (m+n)%2 == 0 {
// 		return float64(median1+median2) / 2.0
// 	}
// 	return float64(median1)
// }

// func max(x, y int) int {
// 	if x > y {
// 		return x
// 	}
// 	return y
// }

// func min(x, y int) int {
// 	if x < y {
// 		return x
// 	}
// 	return y
// }
