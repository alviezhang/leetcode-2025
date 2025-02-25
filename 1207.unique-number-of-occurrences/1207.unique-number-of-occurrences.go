/*
 * @lc app=leetcode.cn id=1207 lang=golang
 *
 * [1207] Unique Number of Occurrences
 */

// @lc code=start
func uniqueOccurrences(arr []int) bool {
	numMap := make(map[int]int)
	for _, num := range arr {
		_, exists := numMap[num]
		if exists {
			numMap[num]++
		} else {
			numMap[num] = 1
		}
	}
	occuMap := make(map[int]struct{})

	for num := range numMap {
		_, exists := occuMap[numMap[num]]
		if exists {
			return false
		} else {
			occuMap[numMap[num]] = struct{}{}
		}
	}
	return true
}

// @lc code=end

