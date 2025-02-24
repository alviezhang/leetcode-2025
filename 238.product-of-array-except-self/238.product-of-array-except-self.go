/*
 * @lc app=leetcode.cn id=238 lang=golang
 *
 * [238] Product of Array Except Self
 */

// @lc code=start
func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	head, tail := 1, 1
	productHead := make([]int, len(nums))
	productTail := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		head = head * nums[i]
		productHead[i] = head
		tail = tail * nums[len(nums)-1-i]
		productTail[len(nums)-1-i] = tail
	}
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			result[i] = productTail[i+1]
		} else if i == len(nums)-1 {
			result[i] = productHead[i-1]
		} else {
			result[i] = productHead[i-1] * productTail[i+1]
		}
	}
	return result
}

// @lc code=end
