/*
 * @lc app=leetcode.cn id=338 lang=golang
 *
 * [338] Counting Bits
 */

// @lc code=start

func count(num int) int {
	c := 0
	for i := 0; i < 18; i++ {
		if 1<<i&num != 0 {
			c++
		}
	}
	return c
}

func countBits(n int) []int {
	ans := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		ans[i] = count(i)
	}
	return ans
}

// @lc code=end

