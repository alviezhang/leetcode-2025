/*
 * @lc app=leetcode.cn id=1137 lang=golang
 *
 * [1137] N-th Tribonacci Number
 */

// @lc code=start

var _tribonacci []int = []int{0, 1, 1}

func tribonacci(n int) int {
	if n < 3 {
		return _tribonacci[n]
	}
	if n < len(_tribonacci) {
		return _tribonacci[n]
	}
	for i := len(_tribonacci); i < n+1; i++ {
		_tribonacci = append(_tribonacci, _tribonacci[i-3]+_tribonacci[i-2]+_tribonacci[i-1])
	}
	return _tribonacci[n]
}

// @lc code=end

