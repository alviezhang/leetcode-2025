/*
 * @lc app=leetcode.cn id=345 lang=golang
 *
 * [345] Reverse Vowels of a String
 */

// @lc code=start

func reverseVowels(s string) string {
	vowels := "aeiouAEIOU"
	new := []byte(s)
	start, end := 0, len(s)-1
	for start < end {
		for !strings.Contains(vowels, string(new[start])) && start < end {
			start++
		}
		for !strings.Contains(vowels, string(new[end])) && start < end {
			end--
		}
		if start < end {
			new[start], new[end] = new[end], new[start]
			start++
			end--
		}
	}
	return string(new)
}

// @lc code=end
