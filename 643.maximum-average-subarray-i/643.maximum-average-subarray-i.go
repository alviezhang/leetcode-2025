/*
 * @lc app=leetcode.cn id=643 lang=golang
 *
 * [643] Maximum Average Subarray I
 */

// @lc code=start
func findMaxAverage(nums []int, k int) float64 {
	maxSum := 0.0
	for i := 0; i < k; i++ {
		maxSum = maxSum + float64(nums[i])
	}

	current := maxSum
	for i := k; i < len(nums); i++ {
		current = current + float64(nums[i]-nums[i-k])
		if current > maxSum {
			maxSum = current
		}
	}
	return maxSum / float64(k)
}

// @lc code=end
