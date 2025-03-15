/*
 * @lc app=leetcode.cn id=452 lang=golang
 *
 * [452] Minimum Number of Arrows to Burst Balloons
 */

// @lc code=start
func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	// fmt.Println(points)

	ans := 0
	lastPos := 0
	for index, point := range points {
		if index == 0 || lastPos < point[0] {
			ans++
			lastPos = point[1]
		}
	}

	return ans
}

// @lc code=end
