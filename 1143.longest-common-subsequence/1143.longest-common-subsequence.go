/*
 * @lc app=leetcode.cn id=1143 lang=golang
 *
 * [1143] Longest Common Subsequence
 */

// @lc code=start
func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text2)+1)
	for i := 0; i <= len(text2); i++ {
		dp[i] = make([]int, len(text1)+1)
	}

	for i := 1; i <= len(text2); i++ {
		for j := 1; j <= len(text1); j++ {
			if text1[j-1] == text2[i-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	// for _, v := range dp {
	// 	fmt.Println(v)
	// }
	return dp[len(text2)][len(text1)]
}

// @lc code=end

