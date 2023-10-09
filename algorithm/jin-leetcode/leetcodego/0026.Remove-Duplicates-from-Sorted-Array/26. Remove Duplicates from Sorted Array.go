package leetcode

// 解法一
func reD1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	last, finder := 0, 0
	for last < len(nums)-1 {
		for nums[finder] == nums[last] {
			finder++
			if finder == len(nums) {
				return last + 1
			}
		}
		nums[last+1] = nums[finder]
		last++
	}
	return last + 1
}
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	last, finder := 0, 0
	for last < len(nums)-1 {
		for nums[finder] == nums[last] {
			finder++
			if finder == len(nums) {
				return last + 1
			}
		}
		nums[last+1] = nums[finder]
		last++
	}
	return last + 1
}

// 解法二
func removeDuplicates1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	length := len(nums)
	lastNum := nums[length-1]
	i := 0
	for i = 0; i < length-1; i++ {
		if nums[i] == lastNum {
			break
		}
		if nums[i+1] == nums[i] {
			removeElement1(nums, i+1, nums[i])
			// fmt.Printf("此时 num = %v length = %v\n", nums, length)
		}
	}
	return i + 1
}

func removeElement1(nums []int, start, val int) int {
	if len(nums) == 0 {
		return 0
	}
	j := start
	for i := start; i < len(nums); i++ {
		if nums[i] != val {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
				j++
			} else {
				j++
			}
		}
	}
	return j
}

// class Solution {
// 	public:
// 		int removeDuplicates(int A[], int n) {

// 			if (n<=1) return n;

// 			int pos=0;
// 			for(int i=0; i<n-1; i++){
// 				if (A[i]!=A[i+1]){
// 					A[++pos] = A[i+1];
// 				}
// 			}
// 			return pos+1;
// 		}
// 	};
func reED(nums []int) int {
	n := len(nums)
	pos := 0
	for i := 0; i < n-1; i++ {
		if nums[i] != nums[i+1] {
			nums[pos+1] = nums[i+1]
			pos++
		}

	}
	return pos + 1
}
