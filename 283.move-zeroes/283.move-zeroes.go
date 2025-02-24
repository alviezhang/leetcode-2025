/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] Move Zeroes
 */

// @lc code=start
func moveZeroes(nums []int) {
	pre := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		}
		if pre != i {
			nums[pre], nums[i] = nums[i], nums[pre]
		}
		pre++
	}
}

// @lc code=end
