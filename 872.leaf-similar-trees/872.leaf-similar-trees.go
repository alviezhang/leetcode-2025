/*
 * @lc app=leetcode.cn id=872 lang=golang
 *
 * [872] Leaf-Similar Trees
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
import "reflect"

func visit(root *TreeNode, leaf *[]int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		*leaf = append(*leaf, root.Val)
		return
	}
	visit(root.Left, leaf)
	visit(root.Right, leaf)
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leaf1 := make([]int, 0)
	leaf2 := make([]int, 0)
	visit(root1, &leaf1)
	visit(root2, &leaf2)
	// fmt.Println(leaf1, leaf2)
	return reflect.DeepEqual(leaf1, leaf2)
}

// @lc code=end
