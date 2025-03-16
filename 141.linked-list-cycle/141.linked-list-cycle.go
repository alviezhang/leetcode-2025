/*
 * @lc app=leetcode.cn id=141 lang=golang
 *
 * [141] Linked List Cycle
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	fast, slow := head.Next.Next, head.Next
	for fast != nil && fast.Next != nil && fast != slow {
		fast = fast.Next.Next
		slow = slow.Next
	}
	if fast == nil || fast.Next == nil {
		return false
	}
	return true
}

// @lc code=end
