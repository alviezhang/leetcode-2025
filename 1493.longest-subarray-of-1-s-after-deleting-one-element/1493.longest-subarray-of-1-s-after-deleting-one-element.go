/*
 * @lc app=leetcode.cn id=1493 lang=golang
 *
 * [1493] Longest Subarray of 1's After Deleting One Element
 */

// @lc code=start
func longestSubarray(nums []int) int {
	// fmt.Println("=================")
	left := 0
	deleted := 0
	ans := 0
	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			deleted++
		}
		for deleted > 1 {
			if nums[left] == 0 {
				deleted--
			}
			left++
		}
		// fmt.Println(nums[left : right+1])
		if right-left+1 == len(nums) {
			ans = max(ans, right-left)
		} else {
			ans = max(ans, right-left+1-deleted)
		}
	}
	return ans
}

// @lc code=end

