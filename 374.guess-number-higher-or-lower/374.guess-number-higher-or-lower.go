/*
 * @lc app=leetcode.cn id=374 lang=golang
 *
 * [374] Guess Number Higher or Lower
 */

// @lc code=start
/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	mid, result := 0, 1
	left, right := 1, n
	for result != 0 {
		mid = (left + right) / 2
		result = guess(mid)
		if result == 1 {
			left = mid + 1
		} else if result == -1 {
			right = mid - 1
		}
	}
	return mid
}

// @lc code=end

