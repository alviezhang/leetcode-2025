/*
 * @lc app=leetcode.cn id=1768 lang=golang
 *
 * [1768] Merge Strings Alternately
 */

// @lc code=start
func mergeAlternately(word1 string, word2 string) string {
	var merged []byte
	var i, j int
	len1 := len(word1)
	len2 := len(word2)
	for i < len1 || j < len2 {
		if i < len1 {
			merged = append(merged, word1[i])
			i++
		}
		if j < len2 {
			merged = append(merged, word2[j])
			j++
		}
	}
	return string(merged)
}

// @lc code=ende
