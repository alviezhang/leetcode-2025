/*
 * @lc app=leetcode.cn id=547 lang=golang
 *
 * [547] Number of Provinces
 */

// @lc code=start
func findCircleNum(isConnected [][]int) int {
	visited := make([]int, len(isConnected))
	queue := make([]int, 0)
	ans := 0

	for i := 0; i < len(visited); i++ {
		if visited[i] != 0 {
			continue
		}

		ans++
		visited[i] = ans
		queue = append(queue, i)

		for len(queue) > 0 {
			current := queue[0]
			queue = queue[1:]
			for j := 0; j < len(visited); j++ {
				if j != current && visited[j] == 0 && isConnected[current][j] == 1 {
					visited[j] = visited[current]
					queue = append(queue, j)
				}
			}
		}
	}
	return ans
}

// @lc code=end

