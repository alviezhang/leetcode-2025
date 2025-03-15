/*
 * @lc app=leetcode.cn id=994 lang=golang
 *
 * [994] Rotting Oranges
 */

// @lc code=start
var directions [4][2]int = [4][2]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func orangesRotting(grid [][]int) int {
	rows := len(grid)
	columns := len(grid[0])

	orangeCount := 0
	queue := make([][2]int, 0)

	visited := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if grid[i][j] == 1 {
				orangeCount++
			} else if grid[i][j] == 2 {
				queue = append(queue, [2]int{i, j})
				orangeCount++
				visited++
			}
		}
	}

	if orangeCount == visited {
		return 0
	}

	for minutes := 1; len(queue) > 0; minutes++ {
		temp := queue
		queue = nil
		for len(temp) > 0 {
			current := temp[0]
			temp = temp[1:]

			for _, d := range directions {
				r := current[0] + d[0]
				c := current[1] + d[1]
				if 0 <= r && r < rows && 0 <= c && c < columns && grid[r][c] == 1 {
					visited++
					grid[r][c] = 2
					queue = append(queue, [2]int{r, c})
				}
			}
		}

		if visited == orangeCount {
			return minutes
		}

	}
	return -1
}

// @lc code=end

