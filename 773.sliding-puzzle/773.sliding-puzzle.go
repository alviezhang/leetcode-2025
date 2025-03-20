/*
 * @lc app=leetcode.cn id=773 lang=golang
 *
 * [773] Sliding Puzzle
 */

// @lc code=start

// 0 1 2
// 3 4 5
var neighbors [6][]int = [6][]int{{1, 3}, {0, 2, 4}, {1, 5}, {0, 4}, {1, 3, 5}, {2, 4}}

func slidingPuzzle(board [][]int) int {
	builder := strings.Builder{}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			builder.WriteByte(byte(int8(board[i][j]) + int8(0x30)))
		}
	}

	visited := map[string]struct{}{}
	queue := []string{builder.String()}

	for steps := 0; len(queue) > 0; steps++ {
		currentQueue := queue
		queue = nil
		for len(currentQueue) > 0 {
			current := currentQueue[0]

			if current == "123450" {
				return steps
			}

			currentQueue = currentQueue[1:]
			visited[current] = struct{}{}

			pos := strings.Index(current, "0")
			for _, direction := range neighbors[pos] {
				newBoard := []byte(current)
				newBoard[pos], newBoard[direction] = newBoard[direction], newBoard[pos]

				newBoardStr := string(newBoard)

				_, exists := visited[newBoardStr]
				if !exists {
					queue = append(queue, newBoardStr)
				}
			}
		}
	}
	return -1
}

// @lc code=end

