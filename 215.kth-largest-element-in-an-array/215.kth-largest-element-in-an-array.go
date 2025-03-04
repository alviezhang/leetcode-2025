/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] Kth Largest Element in an Array
 */

// @lc code=start

func quickSelect(nums []int, k int, left int, right int) int {
	if left == right {
		return nums[k]
	}

	mid := nums[left]

	i := left - 1
	j := right + 1
	for i < j {
		for i++; nums[i] < mid; i++ {
		}
		for j--; mid < nums[j]; j-- {
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	if k <= j {
		return quickSelect(nums, k, left, j)
	} else {
		return quickSelect(nums, k, j+1, right)
	}
}

func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, len(nums)-k, 0, len(nums)-1)
}

// @lc code=end

