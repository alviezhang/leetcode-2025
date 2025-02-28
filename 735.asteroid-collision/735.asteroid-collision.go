/*
 * @lc app=leetcode.cn id=735 lang=golang
 *
 * [735] Asteroid Collision
 */
package main

// @lc code=start
func asteroidCollision(asteroids []int) []int {
	stack := make([]int, len(asteroids))
	i := 0
	for _, value := range asteroids {
		if value > 0 {
			stack[i] = value
			i++
		} else {
			for {
				if i == 0 || stack[i-1] < 0 {
					stack[i] = value
					i++
					break
				}
				if stack[i-1] > -value {
					break
				} else if stack[i-1] == -value {
					i--
					break
				} else {
					i--
				}
			}
		}
	}
	return stack[0:i]
}

// @lc code=end
