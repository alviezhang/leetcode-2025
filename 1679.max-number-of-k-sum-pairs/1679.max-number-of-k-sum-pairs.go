/*
 * @lc app=leetcode.cn id=1679 lang=golang
 *
 * [1679] Max Number of K-Sum Pairs
 */

// @lc code=start

import "sort"

func maxOperations(nums []int, k int) int {
	sort.Ints(nums)
	i, j := 0, len(nums)-1
	count := 0
	// fmt.Println(nums)
	for i < j {
		// fmt.Println(i, j)
		if nums[i]+nums[j] == k {
			count++
			i++
			j--
			continue
		}
		if nums[i] < k-nums[j] {
			i++
		} else {
			j--
		}
	}
	return count
}

// @lc code=end

