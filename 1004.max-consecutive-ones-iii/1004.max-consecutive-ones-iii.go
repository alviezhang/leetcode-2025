/*
 * @lc app=leetcode.cn id=1004 lang=golang
 *
 * [1004] Max Consecutive Ones III
 */
// @lc code=start
func longestOnes(nums []int, k int) int {
	left := 0
	used := 0
	maxCount := 0
	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			used++
		}
		for used > k {
			if nums[left] == 0 {
				used--
			}
			left++
		}
		if right-left+1 > maxCount {
			maxCount = right - left + 1
		}
	}
	return maxCount
}

// @lc code=end
