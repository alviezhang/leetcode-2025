/*
 * @lc app=leetcode.cn id=2095 lang=golang
 *
 * [2095] Delete the Middle Node of a Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *ListNode) *ListNode {
	var pre *ListNode
	fast := head
	slow := head
	count := 0
	for fast != nil {
		if count%2 == 1 {
			pre = slow
			slow = slow.Next
		}
		fast = fast.Next
		count++
	}
	if count == 1 {
		return nil
	}
	// fmt.Println(count, slow)
	pre.Next = slow.Next
	return head
}

// @lc code=end

