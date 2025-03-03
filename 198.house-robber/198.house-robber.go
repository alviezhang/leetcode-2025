/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] House numsber
 */

// @lc code=start
func rob(nums []int) int {
	maxNum := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			maxNum[0] = max(nums[0])
		} else if i == 1 {
			maxNum[1] = max(nums[0], nums[1])
		} else if i == 2 {
			maxNum[2] = max(maxNum[1], nums[2]+maxNum[0])
		} else {
			maxNum[i] = max(maxNum[i-1], nums[i]+maxNum[i-2])
		}
	}
	// fmt.Println(maxNum)
	return maxNum[len(nums)-1]
}

// @lc code=end

