/*
 * @lc app=leetcode.cn id=64 lang=golang
 *
 * [64] Minimum Path Sum
 */

// @lc code=start
func minPathSum(grid [][]int) int {
	rows := len(grid)
	columns := len(grid[0])

	dp := make([][]int, rows)
	for i := 0; i < rows; i++ {
		dp[i] = make([]int, columns)
	}

	for i := rows - 1; i >= 0; i-- {
		for j := columns - 1; j >= 0; j-- {
			if i == rows-1 && j == columns-1 {
				dp[i][j] = grid[i][j]
			} else if i == rows-1 {
				dp[i][j] = dp[i][j+1] + grid[i][j]
			} else if j == columns-1 {
				dp[i][j] = dp[i+1][j] + grid[i][j]
			} else {
				dp[i][j] = min(dp[i+1][j], dp[i][j+1]) + grid[i][j]
			}
		}
	}
	// for i := 0; i < rows; i++ {
	// 	fmt.Println(dp[i])
	// }
	return dp[0][0]
}

// @lc code=end

