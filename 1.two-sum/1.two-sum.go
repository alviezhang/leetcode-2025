/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	mapping := make(map[int]int)
	ans := make([]int, 2)

	for index, num := range nums {
		wanted := target - num
		wantedIndex, exists := mapping[wanted]

		if exists && index != wantedIndex {
			ans[0] = wantedIndex
			ans[1] = index
			break
		}

		mapping[num] = index
	}

	return ans
}

// @lc code=end

