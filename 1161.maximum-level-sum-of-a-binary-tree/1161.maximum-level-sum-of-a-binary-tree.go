/*
 * @lc app=leetcode.cn id=1161 lang=golang
 *
 * [1161] Maximum Level Sum of a Binary Tree
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxLevelSum(root *TreeNode) int {
	var queue []*TreeNode
	queue = append(queue, root)
	maxSum := -1 << 31
	ans := -1
	for i := 1; len(queue) != 0; i++ {
		level := queue
		queue = nil
		sum := 0
		for _, item := range level {
			sum += item.Val
			if item.Left != nil {
				queue = append(queue, item.Left)
			}
			if item.Right != nil {
				queue = append(queue, item.Right)
			}
		}
		if sum > maxSum {
			maxSum = sum
			ans = i
		}
	}
	return ans
}

// @lc code=end

