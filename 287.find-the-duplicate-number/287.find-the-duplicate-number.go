/*
 * @lc app=leetcode.cn id=287 lang=golang
 *
 * [287] Find the Duplicate Number
 */

// @lc code=start
func findDuplicate(nums []int) int {
	slow, fast := 0, 0
	for slow, fast = nums[slow], nums[nums[fast]]; slow != fast; slow, fast = nums[slow], nums[nums[fast]] {
	}
	fast = 0
	for slow != fast {
		fast = nums[fast]
		slow = nums[slow]
	}
	return slow
}

// @lc code=end

