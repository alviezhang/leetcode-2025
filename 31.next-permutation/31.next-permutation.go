/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] Next Permutation
 */

// @lc code=start
func nextPermutation(nums []int) {
	length := len(nums)
	if length <= 1 {
		return
	}

	var i, j int
	for i = length - 1; i > 0 && nums[i-1] >= nums[i]; i-- {
	}
	if i > 0 {
		for j = length - 1; j > i && nums[i-1] >= nums[j]; j-- {

		}
		// fmt.Println(i, j)

		nums[i-1], nums[j] = nums[j], nums[i-1]
	}

	j = length - 1

	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}

}

// @lc code=end
