/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] Letter Combinations of a Phone Number
 */

// @lc code=start
var MAPPING = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func letterCombinations(digits string) []string {
	ans := make([]string, 0)

	for len(digits) > 0 {
		if len(ans) == 0 {
			for _, ch := range MAPPING[digits[0]] {
				ans = append(ans, string(ch))
			}
		} else {
			current := MAPPING[digits[0]]
			newAns := make([]string, len(ans)*len(current))
			for i, s := range ans {
				for j, ch := range current {
					newAns[i*len(current)+j] = s + string(ch)
				}
			}
			ans = newAns
		}
		digits = digits[1:]
	}

	return ans
}

// @lc code=end
