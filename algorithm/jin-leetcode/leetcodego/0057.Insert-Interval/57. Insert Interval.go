package leetcode

import (
	"github.com/halfrost/LeetCode-Go/structures"
)

// Interval define
type Interval = structures.Interval

/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */

func insert(intervals []Interval, newInterval Interval) []Interval {
	res := make([]Interval, 0)
	if len(intervals) == 0 {
		res = append(res, newInterval)
		return res
	}
	curIndex := 0
	for curIndex < len(intervals) && intervals[curIndex].End < newInterval.Start {
		res = append(res, intervals[curIndex])
		curIndex++
	}

	for curIndex < len(intervals) && intervals[curIndex].Start <= newInterval.End {
		newInterval = Interval{Start: min(newInterval.Start, intervals[curIndex].Start), End: max(newInterval.End, intervals[curIndex].End)}
		curIndex++
	}

	res = append(res, newInterval)

	for curIndex < len(intervals) {
		res = append(res, intervals[curIndex])
		curIndex++
	}
	return res
}
func myinsert(intervals []Interval, newInterval Interval) []Interval {
	res := []Interval{}
	if len(intervals) == 0 {
		res = append(res, newInterval)
		return res
	}
	i := 0
	for ; i < len(intervals); i++ {
		if intervals[i].End < newInterval.Start {
			res = append(res, intervals[i])
		} else {
			break
		}
	}
	tempInterval := Interval{Start: newInterval.Start, End: newInterval.End}
	for i < len(intervals) {
		if intervals[i].Start <= newInterval.End {
			if intervals[i].Start < tempInterval.Start {
				tempInterval.Start = intervals[i].Start
			}
			if intervals[i].End > tempInterval.End {
				tempInterval.End = intervals[i].End
			}
			i++
		} else {
			break
		}
	}
	res = append(res, tempInterval)
	for ; i < len(intervals); i++ {
		res = append(res, intervals[i])
	}
	return res
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

func partitionSort(a []Interval, lo, hi int) int {
	pivot := a[hi]
	i := lo - 1
	for j := lo; j < hi; j++ {
		if (a[j].Start < pivot.Start) || (a[j].Start == pivot.Start && a[j].End < pivot.End) {
			i++
			a[j], a[i] = a[i], a[j]
		}
	}
	a[i+1], a[hi] = a[hi], a[i+1]
	return i + 1
}
func quickSort(a []Interval, lo, hi int) {
	if lo >= hi {
		return
	}
	p := partitionSort(a, lo, hi)
	quickSort(a, lo, p-1)
	quickSort(a, p+1, hi)
}
