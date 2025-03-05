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
	if len(digits) == 0 {
		return []string{}
	}

	ans := make([]string, 0)

	var dfs func(int, string)

	dfs = func(i int, s string) {
		if i == len(digits) {
			ans = append(ans, s)
		} else {
			for _, ch := range MAPPING[digits[i]] {
				dfs(i+1, s+string(ch))
			}
		}
	}

	dfs(0, "")

	return ans
}

// @lc code=end
