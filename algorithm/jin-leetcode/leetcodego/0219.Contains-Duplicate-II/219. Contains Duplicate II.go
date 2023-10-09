package leetcode

func containsNearbyDuplicate(nums []int, k int) bool {
	if len(nums) <= 1 {
		return false
	}
	if k <= 0 {
		return false
	}
	record := make(map[int]bool, len(nums))
	for i, n := range nums {
		if _, found := record[n]; found {
			return true
		}
		record[n] = true
		if len(record) == k+1 {
			delete(record, nums[i-k])
		}
	}
	return false
}

func mycontainsNearbyDuplicate(nums []int, k int) bool {
	if len(nums) <= 1 || k <= 0 {
		return false
	}
	record := make(map[int]bool, len(nums))
	for i, n := range nums {
		if _, ok := record[n]; ok {
			return true
		}
		record[n] = true
		if len(record) == k+1 {
			delete(record, nums[i-k])
		}
	}
	return false
}

//	bool containsNearbyDuplicate(vector<int>& nums, int k) {
//		unordered_map<int, int> m;
//		for (int i=0; i<nums.size(); i++) {
//			int n = nums[i];
//			if (m.find(n) != m.end() && i - m[n] <= k) {
//				return true;
//			}
//			m[n] = i;
//		}
//		return false;
//	}
