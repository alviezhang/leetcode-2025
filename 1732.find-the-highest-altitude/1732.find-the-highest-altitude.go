/*
 * @lc app=leetcode.cn id=1732 lang=golang
 *
 * [1732] Find the Highest Altitude
 */

// @lc code=start
func largestAltitude(gain []int) int {
	alti := 0
	ans := 0
	for _, value := range gain {
		alti += value
		ans = max(ans, alti)
	}
	return ans
}

// @lc code=end

