/*
 * @lc app=leetcode.cn id=334 lang=golang
 *
 * [334] Increasing Triplet Subsequence
 */

// @lc code=start
func increasingTriplet(nums []int) bool {
	maxNums := make([]int, len(nums))
	maxNums[len(nums)-1] = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] > maxNums[i+1] {
			maxNums[i] = nums[i]
		} else {
			maxNums[i] = maxNums[i+1]
		}
	}

	minNums := make([]int, len(nums))
	minNums[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < minNums[i-1] {
			minNums[i] = nums[i]
		} else {
			minNums[i] = minNums[i-1]
		}
	}

	// fmt.Println(nums)
	// fmt.Println(minNums)
	// fmt.Println(maxNums)

	for i := 1; i < len(nums)-1; i++ {
		if minNums[i-1] < nums[i] && nums[i] < maxNums[i+1] {
			return true
		}
	}
	return false
}

// @lc code=end
