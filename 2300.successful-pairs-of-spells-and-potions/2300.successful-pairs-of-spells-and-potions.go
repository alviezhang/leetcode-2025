/*
 * @lc app=leetcode.cn id=2300 lang=golang
 *
 * [2300] Successful Pairs of Spells and Potions
 */

// @lc code=start
func successfulPairs(spells []int, potions []int, success int64) []int {
	ans := make([]int, len(spells))
	sort.Sort(sort.IntSlice(potions))

	for i := 0; i < len(spells); i++ {
		minimumPotion := int(success / int64(spells[i]))
		if success%int64(spells[i]) > 0 {
			minimumPotion++
		}

		left := 0
		right := len(potions) - 1
		for left <= right {
			mid := (left + right) / 2
			if mid == 0 && potions[mid] >= minimumPotion || potions[mid] >= minimumPotion && potions[mid-1] < minimumPotion {
				ans[i] = len(potions) - mid
				break
			} else if potions[mid] >= minimumPotion && potions[mid-1] >= minimumPotion {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}

	}
	return ans
}

// @lc code=end

