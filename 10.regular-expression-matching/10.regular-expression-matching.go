/*
 * @lc app=leetcode.cn id=10 lang=golang
 *
 * [10] Regular Expression Matching
 */

// @lc code=start
func isMatch(s string, p string) bool {
	// fmt.Println("match:", s, p)
	var target byte
	si, pi := 0, 0
	hasStar := false

	for si <= len(s) && pi < len(p) {
		target = p[pi]
		if pi+1 < len(p) && p[pi+1] == '*' {
			hasStar = true
			pi += 2
		} else {
			pi++
		}

		if !hasStar {
			if si < len(s) && (target == '.' || target == s[si]) {
				target = 0
				si++
			} else {
				return false
			}
		} else {
			for si <= len(s) {
				if isMatch(s[si:], p[pi:]) {
					return true
				} else {
					if si < len(s) && (target == '.' || target == s[si]) {
						si++
					} else {
						return false
					}
				}
			}
		}
	}
	// fmt.Println(si, pi, target, hasStar)
	if si == len(s) && pi == len(p) && (target == 0 && hasStar == false || hasStar == true) {
		return true
	}
	return false
}

// @lc code=end
