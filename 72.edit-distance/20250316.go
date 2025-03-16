/*
 * @lc app=leetcode.cn id=72 lang=golang
 *
 * [72] Edit Distance
 */

// @lc code=start
func minDistance(word1 string, word2 string) int {
	mem := make([][]int, len(word1))
	for i := 0; i < len(mem); i++ {
		mem[i] = make([]int, len(word2))
		for j := 0; j < len(mem[i]); j++ {
			mem[i][j] = -1
		}
	}

	var dp func(int, int) int
	dp = func(i int, j int) int {
		if i < 0 {
			return j + 1
		}
		if j < 0 {
			return i + 1
		}
		if mem[i][j] == -1 {
			if word1[i] == word2[j] {
				mem[i][j] = dp(i-1, j-1)
			} else {
				mem[i][j] = min(dp(i-1, j), dp(i, j-1), dp(i-1, j-1)) + 1
			}
		}

		return mem[i][j]
	}
	return dp(len(word1)-1, len(word2)-1)
}

// @lc code=end
