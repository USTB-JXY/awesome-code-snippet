package leetcode

import "sort"

// 解法一
func hIndex(citations []int) int {
	n := len(citations)
	buckets := make([]int, n+1)
	for _, c := range citations {
		if c >= n {
			buckets[n]++
		} else {
			buckets[c]++
		}
	}
	count := 0
	for i := n; i >= 0; i-- {
		count += buckets[i]
		if count >= i {
			return i
		}
	}
	return 0
}

// 解法二
func hIndex1(citations []int) int {
	quickSort164(citations, 0, len(citations)-1)
	hIndex := 0
	for i := len(citations) - 1; i >= 0; i-- {
		if citations[i] >= len(citations)-i {
			hIndex++
		} else {
			break
		}
	}
	return hIndex
}

func quickSort164(a []int, lo, hi int) {
	if lo >= hi {
		return
	}
	p := partition164(a, lo, hi)
	quickSort164(a, lo, p-1)
	quickSort164(a, p+1, hi)
}

func partition164(a []int, lo, hi int) int {
	pivot := a[hi]
	i := lo - 1
	for j := lo; j < hi; j++ {
		if a[j] < pivot {
			i++
			a[j], a[i] = a[i], a[j]
		}
	}
	a[i+1], a[hi] = a[hi], a[i+1]
	return i + 1
}

// h指数由J.E. Hirsch于2005年提出，发表在《美国国家科学院院刊》上。[i] h指数是一种定量指标，基于对出版物和引用的出版物数据分析，以提供“对科学家累积研究贡献的重要性，重要性和广泛影响的估计”。[ii] 根据赫希的说法，h指数被定义为：“如果科学家的Np论文中有h篇至少有h次引用，而其他（Np – h）论文每篇都有≤h次引用，那么他就有索引h
//
// {                                     1 2 3 4
// 	para274{[]int{3, 6, 9, 1}},          1 3 6 9
// 	ans274{3},                           0 1 2 3
// },                                    4 3 2 1

// sort.Ints(s)
func kindex(nums []int) int {
	sort.Ints(nums)
	n := len(nums)

	for i := 0; i < n; i++ {
		if nums[i] >= n-i {
			return nums[i]
		}
	}
	return 0
}
