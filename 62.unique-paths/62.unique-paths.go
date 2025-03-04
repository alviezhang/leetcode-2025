/*
 * @lc app=leetcode.cn id=62 lang=golang
 *
 * [62] Unique Paths
 */
// @lc code=start
func uniquePaths(m int, n int) int {
	s := make([][]int, m+1)
	for i := 0; i < len(s); i++ {
		s[i] = make([]int, n+1)
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				s[i][j] = 1
				continue
			}
			s[i][j] = s[i+1][j] + s[i][j+1]
		}
	}
	// for i := 0; i < len(s); i++ {
	// 	fmt.Println(s[i])
	// }
	return s[0][0]
}

// @lc code=end
