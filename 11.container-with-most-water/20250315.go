/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] Container With Most Water
 */

// @lc code=start

// type MaximumInfo struct {
// 	pos   int
// 	value int
// }

func maxArea(height []int) int {
	i, j := 0, len(height)-1
	ans := -1
	for i < j {
		current := min(height[i], height[j]) * (j - i)
		if current > ans {
			ans = current
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return ans
}

// @lc code=end
