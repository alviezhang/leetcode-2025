/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] Trapping Rain Water
 */

// @lc code=start
func trap(height []int) int {
	ans := 0
	left := 0
	for i := 1; i < len(height); i++ {
		if height[left] <= height[i] {
			for j := left + 1; j < i; j++ {
				ans += height[left] - height[j]
			}
			left = i
		}
	}

	right := len(height) - 1
	for i := len(height) - 2; i >= left; i-- {
		if height[i] >= height[right] {
			for j := right - 1; j > i; j-- {
				ans += height[right] - height[j]
			}
			right = i
		}
	}
	return ans
}

// @lc code=end
