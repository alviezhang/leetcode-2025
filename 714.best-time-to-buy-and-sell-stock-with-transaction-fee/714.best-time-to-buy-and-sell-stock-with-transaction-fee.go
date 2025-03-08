/*
 * @lc app=leetcode.cn id=714 lang=golang
 *
 * [714] Best Time to Buy and Sell Stock with Transaction Fee
 */

// @lc code=start
func maxProfit(prices []int, fee int) int {
	length := len(prices)
	dp := make([][]int, length+1)
	for i := 0; i < length+1; i++ {
		dp[i] = make([]int, 2)
	}

	dp[0][0] = 0
	dp[0][1] = -1 << 31

	for i := 1; i < length+1; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i-1]-fee)
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i-1])
	}
	return dp[length][0]
}

// @lc code=end

