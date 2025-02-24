/*
 * @lc app=leetcode.cn id=392 lang=golang
 *
 * [392] Is Subsequence
 */

// @lc code=start
func isSubsequence(s string, t string) bool {
	si, ti := 0, 0
	for si < len(s) && ti < len(t) {
		if s[si] == t[ti] {
			si++
		}
		ti++
	}
	if si == len(s) {
		return true
	}
	return false
}

// @lc code=end

