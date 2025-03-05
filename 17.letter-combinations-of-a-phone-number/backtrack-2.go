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
	path := make([]byte, len(digits))

	var dfs func(int)

	dfs = func(i int) {
		if i == len(digits) {
			ans = append(ans, string(path))
		} else {
			for _, ch := range MAPPING[digits[i]] {
				path[i] = byte(ch)
				dfs(i + 1)
			}
		}
	}

	dfs(0)

	return ans
}

// @lc code=end
