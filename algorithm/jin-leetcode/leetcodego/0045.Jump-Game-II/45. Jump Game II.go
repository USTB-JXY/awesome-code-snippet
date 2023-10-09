package leetcode

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	needChoose, canReach, step := 0, 0, 0
	for i, x := range nums {
		if i+x > canReach {
			canReach = i + x
			if canReach >= len(nums)-1 {
				return step + 1
			}
		}
		if i == needChoose {
			needChoose = canReach
			step++
		}
	}
	return step
}

// //Acutally, using the Greedy algorithm can have the answer
// int jump(int A[], int n) {
//     if (n<=1) return 0;

//     int steps = 0;
//     int coverPos = 0;
//     for (int i=0; i<=coverPos&& i<n; ){
//         if (A[i]==0) return -1;
//         if(coverPos < A[i]+i){
//             coverPos = A[i]+i;
//             steps++;
//         }
//         if (coverPos >= n-1){
//             return steps;
//         }
//         //Greedy: find the next place which can have biggest distance
//         int nextPos=0;
//         int maxDistance=0;
//         for(int j=i+1; j<=coverPos; j++){
//             if ( A[j]+j > maxDistance ) {
//                 maxDistance = A[j]+j;
//                 nextPos = j;
//             }
//         }
//         i = nextPos;
//     }
//     return steps;
// }
func cjump2(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	step, cpos := 0, 0
	for i := 0; i < len(nums) && i <= cpos; {
		if nums[i] == 0 && i == cpos {
			return -1
		}
		if nums[i]+i > cpos {
			cpos = nums[i] + i
			step++
		}
		if cpos >= len(nums)-1 {
			return step
		}
		nextPos, Maxpos := 0, 0
		for j := i + 1; j <= cpos; j++ {
			if nums[j]+j > Maxpos {
				nextPos = j
				Maxpos = nums[j] + j
			}
		}
		i = nextPos

	}
	return step

}
