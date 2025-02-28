/*
 * @lc app=leetcode.cn id=649 lang=golang
 *
 * [649] Dota2 Senate
 */

// @lc code=start
func predictPartyVictory(senate string) string {
	radiant := make([]int, 0, len(senate))
	dire := make([]int, 0, len(senate))
	for index, value := range senate {
		if value == 'D' {
			dire = append(dire, index)
		} else {
			radiant = append(radiant, index)
		}
	}
	for len(radiant) != 0 && len(dire) != 0 {
		if radiant[0] < dire[0] {
			radiant = append(radiant, radiant[0]+len(senate))[1:]
			dire = dire[1:]
		} else {
			radiant = radiant[1:]
			dire = append(dire, dire[0]+len(senate))[1:]
		}
		// fmt.Println(radiant, dire)
	}
	if len(dire) == 0 {
		return "Radiant"
	} else {
		return "Dire"
	}
}

// @lc code=end
