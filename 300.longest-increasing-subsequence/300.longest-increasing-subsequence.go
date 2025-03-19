/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] Longest Increasing Subsequence
 */

// @lc code=start
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 1

	ans := 1

	for i := 1; i < len(nums); i++ {
		maxLis := 0
		for j := 0; j < i; j++ {
			if dp[j] > maxLis && nums[j] < nums[i] {
				maxLis = dp[j]
			}
		}
		dp[i] = maxLis + 1
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	// fmt.Println(dp)
	return ans
}

// @lc code=end

