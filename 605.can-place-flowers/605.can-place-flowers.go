/*
 * @lc app=leetcode.cn id=605 lang=golang
 *
 * [605] Can Place Flowers
 */

package main

// @lc code=start
func canPlaceFlowers(flowerbed []int, n int) bool {
	numOfFlowers := len(flowerbed)
	numOfPlaces := 0
	startPos := -2
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 1 {
			numOfPlaces += (i - startPos - 2) / 2
			startPos = i
		} else if i == numOfFlowers-1 {
			numOfPlaces += (i - startPos) / 2
		}
	}
	return numOfPlaces >= n
}

// @lc code=end
