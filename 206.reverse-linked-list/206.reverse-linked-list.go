/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] Reverse Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var previous *ListNode
	for head != nil {
		t := head.Next
		head.Next = previous
		previous = head
		head = t
	}
	return previous
}

// @lc code=end

