/*
 * @lc app=leetcode.cn id=2215 lang=golang
 *
 * [2215] Find the Difference of Two Arrays
 */

package main

// @lc code=start
func findDifference(nums1 []int, nums2 []int) [][]int {
	set1 := make(map[int]struct{})
	set2 := make(map[int]struct{})
	for _, value := range nums1 {
		_, exist := set1[value]
		if !exist {
			set1[value] = struct{}{}
		}
	}
	for _, value := range nums2 {
		_, exist := set2[value]
		if !exist {
			set2[value] = struct{}{}
		}
	}

	ans := make([][]int, 2)
	ans[0] = make([]int, 0)
	ans[1] = make([]int, 0)

	for num := range set1 {
		_, exists := set2[num]
		if !exists {
			ans[0] = append(ans[0], num)
		}
	}
	for num := range set2 {
		_, exists := set1[num]
		if !exists {
			ans[1] = append(ans[1], num)
		}
	}
	return ans
}

// @lc code=endj
