/*
 * @lc app=leetcode.cn id=75 lang=golang
 *
 * [75] Sort Colors
 */

// @lc code=start
func sortColors(nums []int) {
	count := [3]int{}
	for _, color := range nums {
		count[color]++
	}
	color := 0
	for i := 0; i < len(nums); i++ {
		for count[color] == 0 && color < 2 {
			color++
		}
		nums[i] = color
		count[color]--
	}
}

// @lc code=end

