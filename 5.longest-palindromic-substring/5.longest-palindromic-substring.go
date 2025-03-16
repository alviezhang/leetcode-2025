/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 */

// @lc code=start
func longestPalindrome(s string) string {
	dp := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]int, len(s))
	}

	maxAns := 0
	ans := [2]int{}

	for j := 0; j < len(s); j++ {
		for i := j; i >= 0; i-- {
			if i == j {
				dp[i][j] = 1
			} else if i == j-1 && s[i] == s[j] {
				dp[i][j] = 2
			} else if dp[i+1][j-1] > 0 && s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			}

			if dp[i][j] > maxAns {
				maxAns = dp[i][j]
				ans[0], ans[1] = i, j
			}
		}
	}
	// for i := 0; i < len(s); i++ {
	// 	fmt.Println(dp[i])
	// }
	// fmt.Println(maxAns)
	return s[ans[0] : ans[1]+1]
}

// @lc code=end
