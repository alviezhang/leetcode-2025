/*
 * @lc app=leetcode.cn id=1456 lang=golang
 *
 * [1456] Maximum Number of Vowels in a Substring of Given Length
 */

// @lc code=start

func isVowel(char byte) bool {
	switch char {
	case 'a':
		return true
	case 'e':
		return true
	case 'i':
		return true
	case 'o':
		return true
	case 'u':
		return true
	default:
		return false
	}
}

func maxVowels(s string, k int) int {
	vowels := 0
	for i := 0; i < k; i++ {
		if isVowel(s[i]) {
			vowels++
		}
	}

	maxCount := vowels
	for i := k; i < len(s); i++ {
		if isVowel(s[i-k]) {
			vowels--
		}
		if isVowel(s[i]) {
			vowels++
		}
		if vowels > maxCount {
			maxCount = vowels
		}
	}
	return maxCount
}

// @lc code=end

