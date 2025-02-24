/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] Container With Most Water
 */

// @lc code=start

// type MaximumInfo struct {
// 	pos   int
// 	value int
// }

func maxArea(height []int) int {
	// leftMaximum := make([]MaximumInfo, len(height))
	// leftMaximum[0] = MaximumInfo{0, height[0]}

	// for i := 1; i < len(height); i++ {
	// 	if height[i] > leftMaximum[i-1].value {
	// 		leftMaximum[i].pos = i
	// 		leftMaximum[i].value = height[i]
	// 	} else {
	// 		leftMaximum[i].pos = leftMaximum[i-1].pos
	// 		leftMaximum[i].value = leftMaximum[i-1].value
	// 	}
	// }

	// rightMaximum := make([]MaximumInfo, len(height))
	// rightMaximum[len(height)-1] = MaximumInfo{len(height) - 1, height[len(height)-1]}
	// for i := len(height) - 2; i >= 0; i-- {
	// 	if height[i] > rightMaximum[i+1].value {
	// 		rightMaximum[i].pos = i
	// 		rightMaximum[i].value = height[i]
	// 	} else {
	// 		rightMaximum[i].pos = rightMaximum[i+1].pos
	// 		rightMaximum[i].value = rightMaximum[i+1].value
	// 	}
	// }
	// fmt.Println(leftMaximum)
	// fmt.Println(rightMaximum)

	maximum := -1
	i, j := 0, len(height)-1
	for i < j {
		cur := min(height[i], height[j]) * (j - i)
		if cur > maximum {
			maximum = cur
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return maximum
}

// @lc code=end
