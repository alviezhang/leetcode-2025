/*
 * @lc app=leetcode.cn id=746 lang=golang
 *
 * [746] Min Cost Climbing Stairs
 */

// @lc code=start
func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	minCost := make([]int, n+1)
	for i := 0; i <= n; i++ {
		if i < 2 {
			minCost[i] = 0
		} else {
			minCost[i] = min(cost[i-1]+minCost[i-1], cost[i-2]+minCost[i-2])
		}
	}
	return minCost[n]
}

// @lc code=end

