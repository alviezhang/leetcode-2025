/*
 * @lc app=leetcode.cn id=1431 lang=golang
 *
 * [1431] Kids With the Greatest Number of Candies
 */

// @lc code=start
func kidsWithCandies(candies []int, extraCandies int) []bool {
	// find the maximum candy
	max := candies[0]
	for i := 1; i < len(candies); i++ {
		if candies[i] > max {
			max = candies[i]
		}
	}

	var isGreatest []bool = make([]bool, len(candies))
	for i := 0; i < len(candies); i++ {
		if extraCandies+candies[i] >= max {
			isGreatest[i] = true
		}
	}
	return isGreatest
}

// @lc code=end

