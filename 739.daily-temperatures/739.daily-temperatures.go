/*
 * @lc app=leetcode.cn id=739 lang=golang
 *
 * [739] Daily Temperatures
 */

// @lc code=start
func dailyTemperatures(temperatures []int) []int {
	stack := make([]int, 0)
	ans := make([]int, len(temperatures))
	for i, temp := range temperatures {
		for len(stack) > 0 && temp > temperatures[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[top] = i - top
		}
		stack = append(stack, i)
	}
	return ans
}

// @lc code=end

