/*
 * @lc app=leetcode.cn id=2390 lang=golang
 *
 * [2390] Removing Stars From a String
 */

// @lc code=start
func removeStars(s string) string {
	charArray := make([]rune, len(s))
	point := 0
	for _, ch := range s {
		if ch == '*' {
			point--
		} else {
			charArray[point] = ch
			point++
		}
	}
	return string(charArray[0:point])
}

// @lc code=end
