/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] Majority Element
 */

// @lc code=start
func majorityElement(nums []int) int {
	ans := nums[0]
	count := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == ans {
			count++
		} else if count == 1 && nums[i] != ans {
			ans = nums[i]
			count = 1
		} else {
			count--
		}
	}
	return ans
}

// @lc code=end

