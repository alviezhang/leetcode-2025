/*
 * @lc app=leetcode.cn id=724 lang=golang
 *
 * [724] Find Pivot Index
 */

// @lc code=start
func pivotIndex(nums []int) int {
	if len(nums) == 1 {
		return 0
	} else if len(nums) == 2 {
		if nums[1] == 0 {
			return 0
		} else if nums[0] == 0 {
			return 1
		}
		return -1
	}

	leftSum := 0
	rightSum := 0
	for i := len(nums) - 1; i >= 0; i-- {
		rightSum += nums[i]
	}
	for i := 0; i < len(nums); i++ {
		rightSum -= nums[i]
		if leftSum == rightSum {
			return i
		}
		leftSum += nums[i]
	}
	return -1
}

// @lc code=end
