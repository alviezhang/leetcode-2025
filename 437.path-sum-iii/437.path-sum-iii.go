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
func getSum(root *TreeNode, targetSum int, sumArray []int) int {
	if root == nil {
		return 0
	}

	ans := 0

	sumArray = append(sumArray, 0)

	for i := 0; i < len(sumArray); i++ {
		sumArray[i] += root.Val
		if sumArray[i] == targetSum {
			ans++
		}
	}

	left := make([]int, len(sumArray))
	copy(left, sumArray)
	right := make([]int, len(sumArray))
	copy(right, sumArray)

	return ans + getSum(root.Left, targetSum, left) + getSum(root.Right, targetSum, right)
}

func pathSum(root *TreeNode, targetSum int) int {
	return getSum(root, targetSum, []int{})
}

// @lc code=end
