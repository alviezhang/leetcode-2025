/*
 * @lc app=leetcode.cn id=199 lang=golang
 *
 * [199] Binary Tree Right Side View
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
type Queue struct {
	Node  *TreeNode
	Level int
}

func rightSideView(root *TreeNode) []int {
	view := make([]int, 0)
	if root == nil {
		return view
	}

	queue := make([]*Queue, 0)

	queue = append(queue, &Queue{Node: root, Level: 1})
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]

		if len(view) < top.Level {
			view = append(view, top.Node.Val)
		}

		if top.Node.Right != nil {
			queue = append(queue, &Queue{Node: top.Node.Right, Level: top.Level + 1})
		}
		if top.Node.Left != nil {
			queue = append(queue, &Queue{Node: top.Node.Left, Level: top.Level + 1})
		}
	}
	// fmt.Println(view)
	return view
}

// @lc code=end

