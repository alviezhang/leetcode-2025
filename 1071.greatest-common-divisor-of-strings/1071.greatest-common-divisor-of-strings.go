/*
 * @lc app=leetcode.cn id=1071 lang=golang
 *
 * [1071] Greatest Common Divisor of Strings
 */

// @lc code=start
import "strings"

func isDivide(str1 string, str2 string) bool {
	len1 := len(str1)
	i := 0
	for i < len1 {
		if !strings.HasPrefix(str1[i:], str2) {
			return false
		}
		i += len(str2)
	}
	if i == len1 {
		return true
	} else {
		return false
	}
}

func gcdOfStrings(str1 string, str2 string) string {
	len1 := len(str1)
	len2 := len(str2)
	i := 0
	found := -1
	for i < len1 && i < len2 {
		if str1[i] == str2[i] {
			if isDivide(str1, str1[:i]) && isDivide(str2, str2[:i]) {
				found = i
			}
		} else {
			break
		}
	}
	if found == -1 {
		return ""
	} else {
		return str1[:found]
	}
}

// @lc code=end
