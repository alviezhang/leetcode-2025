/*
 * @lc app=leetcode.cn id=790 lang=golang
 *
 * [790] Domino and Tromino Tiling
 */

// @lc code=start

var dp [][4]int = [][4]int{{0, 0, 0, 0}, {1, 1, 0, 0}}

const MOD = 1e9 + 7

func numTilings(n int) int {
	if n >= len(dp) {
		for i := len(dp); i <= n; i++ {
			var current [4]int
			current[0] = dp[i-1][1]
			current[1] = (dp[i-1][0] + dp[i-1][1] + dp[i-1][2] + dp[i-1][3]) % MOD
			current[2] = (dp[i-1][0] + dp[i-1][3]) % MOD
			current[3] = (dp[i-1][0] + dp[i-1][2]) % MOD
			dp = append(dp, current)
		}
	}
	return dp[n][1]
}

// @lc code=end

