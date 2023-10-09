package leetcode

import "fmt"

func removeDuplicates(nums []int) int {
	slow := 0
	for fast, v := range nums {
		if fast < 2 || nums[slow-2] != v {
			nums[slow] = v
			slow++
		}
	}
	return slow
}

//参考耗子叔原地修改
func reD2(nums []int) int {
	pos, cnt := 0, 1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != nums[i+1] {
			nums[pos+1] = nums[i+1]
			pos++
			cnt = 1
		} else {
			if (cnt) < 2 {
				nums[pos+1] = nums[i+1]
				pos++
				cnt++
			}
		}
		fmt.Println(pos)
		fmt.Println(nums)
	}
	fmt.Println("jin    ")
	return pos + 1
}

// int removeDuplicates(int A[], int n) {
//     if (n<=2) return n;
//     int pos=0;
//     int cnt=1;
//     for (int i=1; i<n; i++){
//         if (A[i] == A[pos]){
//             cnt++;
//             if (cnt==2){
//                 A[++pos] = A[i];
//             }
//         }else{
//             cnt=1;
//             A[++pos] = A[i];

//         }
//     }
//     return pos+1;
// }

func reD3(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	pos, cnt := 0, 1
	for i := 0; i < len(nums); i++ {
		if nums[i] == nums[pos] {
			cnt++
			if cnt == 2 {
				nums[pos+1] = nums[i]
				pos++
			}
		} else {
			cnt = 1
			nums[pos+1] = nums[i]
			pos++
		}
		fmt.Println(nums, "----", pos)
	}

	return pos + 1
}
