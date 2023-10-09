package leetcode

import (
	"github.com/halfrost/LeetCode-Go/structures"
)

// ListNode define
type ListNode = structures.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	newHead := &ListNode{Val: 0, Next: head}
	len := 0
	cur := newHead
	for cur.Next != nil {
		len++
		cur = cur.Next
	}
	if (k % len) == 0 {
		return head
	}
	cur.Next = head
	cur = newHead
	for i := len - k%len; i > 0; i-- {
		cur = cur.Next
	}
	res := &ListNode{Val: 0, Next: cur.Next}
	cur.Next = nil
	return res.Next
}

func myrotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	cur, listlenth := head, 0
	for cur != nil {
		listlenth++
		cur = cur.Next
	}
	dummy := &ListNode{Next: head}
	slow, fast := head, head
	for k = k % listlenth; k > 0; k-- {
		fast = fast.Next
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	fast.Next = dummy.Next
	dummy.Next = slow.Next
	slow.Next = nil
	return dummy.Next
}
