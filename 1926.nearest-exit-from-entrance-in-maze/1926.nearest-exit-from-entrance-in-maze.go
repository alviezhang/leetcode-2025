/*
 * @lc app=leetcode.cn id=1926 lang=golang
 *
 * [1926] Nearest Exit from Entrance in Maze
 */

// @lc code=start
var directions [4][2]int = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func nearestExit(maze [][]byte, entrance []int) int {
	rows := len(maze)
	columns := len(maze[0])

	queue := make([][2]int, 0)
	queue = append(queue, [2]int{entrance[0], entrance[1]})
	maze[entrance[0]][entrance[1]] = '+'

	for ans := 1; len(queue) != 0; ans++ {
		tmp := queue
		queue = nil
		for _, item := range tmp {
			for _, d := range directions {
				r := item[0] + d[0]
				c := item[1] + d[1]
				if 0 <= r && r < rows && 0 <= c && c < columns && maze[r][c] == '.' {
					// fmt.Printf("%d, %d, %c\n", r, c, maze[r][c])
					if r == 0 || c == 0 || r == rows-1 || c == columns-1 {
						return ans
					}
					maze[r][c] = '+'
					queue = append(queue, [2]int{r, c})
				}
			}
		}
	}
	return -1
}

// @lc code=end
