/*
 * @lc app=leetcode.cn id=2352 lang=golang
 *
 * [2352] Equal Row and Column Pairs
 */

// @lc code=start
func equalPairs(grid [][]int) int {
	columns := make(map[string]int)
	rows := make(map[string]int)
	ans := 0
	keyArray := make([]string, len(grid))
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			keyArray[j] = string(grid[i][j])
		}
		rowKey := strings.Join(keyArray, ",")
		_, rExists := rows[rowKey]
		if rExists {
			rows[rowKey]++
		} else {
			rows[rowKey] = 1
		}

		for j := 0; j < len(grid[i]); j++ {
			keyArray[j] = string(grid[j][i])
		}
		columnKey := strings.Join(keyArray, ",")
		_, cExists := columns[columnKey]
		if cExists {
			columns[columnKey]++
		} else {
			columns[columnKey] = 1
		}
	}
	for key := range rows {
		_, exists := columns[key]
		if exists {
			ans += rows[key] * columns[key]
		}
	}
	return ans
}

// @lc code=end

