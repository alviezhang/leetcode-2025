/*
 * @lc app=leetcode.cn id=72 lang=golang
 *
 * [72] Edit Distance
 */

// @lc code=start
func minDistance(word1 string, word2 string) int {
	if word1 == "" || word2 == "" {
		return max(len(word1), len(word2))
	}
	state := make([][]int, len(word2))
	for i := 0; i < len(state); i++ {
		state[i] = make([]int, len(word1))
	}
	for j := 0; j < len(word2); j++ {
		for i := 0; i < len(word1); i++ {
			current := 1
			if word1[i] == word2[j] {
				current = 0
			}
			if i == 0 && j == 0 {
				state[j][i] = current
			} else if i == 0 && j != 0 {
				state[j][i] = min(j+current, state[j-1][i]+1)
			} else if i != 0 && j == 0 {
				state[j][i] = min(i+current, state[j][i-1]+1)
			} else {
				state[j][i] = min(state[j-1][i-1]+current, state[j-1][i]+1, state[j][i-1]+1)
			}
		}
	}
	return state[len(word2)-1][len(word1)-1]
}

// @lc code=end
