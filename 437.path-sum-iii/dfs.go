/*
 * @lc app=leetcode.cn id=437 lang=golang
 *
 * [437] Path Sum III
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
func pathSum(root *TreeNode, targetSum int) int {
	ans := 0
	cnt := map[int]int{0: 1}
	var dfs func(*TreeNode, int)

	dfs = func(node *TreeNode, sum int) {
		if node == nil {
			return
		}

		sum += node.Val
		ans += cnt[sum-targetSum]
		cnt[sum]++
		dfs(node.Left, sum)
		dfs(node.Right, sum)
		cnt[sum]--
	}
	dfs(root, 0)
	return ans
}

// @lc code=end
