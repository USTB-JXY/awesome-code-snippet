package leetcode

// 解法一 时间复杂度 O(n) 空间复杂度 O(1)
func majorityElement(nums []int) int {
	res, count := nums[0], 0
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			res, count = nums[i], 1
		} else {
			if nums[i] == res {
				count++
			} else {
				count--
			}
		}
	}
	return res
}

// 解法二 时间复杂度 O(n) 空间复杂度 O(n)
func majorityElement1(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
		if m[v] > len(nums)/2 {
			return v
		}
	}
	return 0
}

// int majorityElement(vector<int> &num) {
//     int majority;
//     int cnt = 0;
//     for(int i=0; i<num.size(); i++){
//         if ( cnt ==0 ){
//             majority = num[i];
//             cnt++;
//         }else{
//             majority == num[i] ? cnt++ : cnt --;
//             if (cnt > num.size()/2) return majority;
//         }
//     }
//     return majority;

// }

func mj(nums []int) int {
	maj, cnt := nums[0], 1
	for i := 0; i < len(nums); i++ {
		if cnt == 0 {
			maj, cnt = nums[i], 1
		} else {
			if maj == nums[i] {
				cnt++
				if cnt > maj/2 {
					return maj
				}
			} else {
				cnt--
			}
		}

	}
	return maj
}
