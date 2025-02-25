/*
 * @lc app=leetcode.cn id=1657 lang=golang
 *
 * [1657] Determine if Two Strings Are Close
 */

// @lc code=start
func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	wordMap1 := make(map[rune]int)
	for _, char := range word1 {
		_, exists := wordMap1[char]
		if exists {
			wordMap1[char]++
		} else {
			wordMap1[char] = 1
		}
	}
	wordMap2 := make(map[rune]int)
	for _, char := range word2 {
		_, exists := wordMap2[char]
		if exists {
			wordMap2[char]++
		} else {
			wordMap2[char] = 1
		}
	}
	if len(wordMap1) != len(wordMap2) {
		return false
	}

	occuMap1 := make(map[int]int)
	for char := range wordMap1 {
		_, charExists := wordMap2[char]
		if !charExists {
			return false
		}

		_, exists := occuMap1[wordMap1[char]]
		if exists {
			occuMap1[wordMap1[char]]++
		} else {
			occuMap1[wordMap1[char]] = 1
		}
	}
	occuMap2 := make(map[int]int)
	for char := range wordMap2 {
		_, charExists := wordMap1[char]
		if !charExists {
			return false
		}

		_, exists := occuMap2[wordMap2[char]]
		if exists {
			occuMap2[wordMap2[char]]++
		} else {
			occuMap2[wordMap2[char]] = 1
		}
	}
	// fmt.Println(occuMap1, occuMap2)
	for occu := range occuMap1 {
		_, exists := occuMap2[occu]
		if !exists || occuMap1[occu] != occuMap2[occu] {
			return false
		} else {
			delete(occuMap2, occu)
		}
	}
	if len(occuMap2) == 0 {
		return true
	} else {
		return false
	}
}

// @lc code=end

