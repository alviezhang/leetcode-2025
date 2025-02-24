/*
 * @lc app=leetcode.cn id=151 lang=golang
 *
 * [151] Reverse Words in a String
 */

package main

import "strings"

// @lc code=start
func reverseWords(s string) string {
	var words []string
	i := len(s) - 1
	inWord := false
	word_end := i + 1
	for i >= 0 {
		if s[i] == ' ' {
			if inWord {
				words = append(words, s[i+1:word_end])
			}
			inWord = false
			word_end = i
		} else {
			if !inWord {
				inWord = true
				word_end = i + 1
			}
		}
		i--
	}
	if inWord && word_end != 0 {
		words = append(words, s[0:word_end])
	}
	return strings.Join(words, " ")
}

// @lc code=end
