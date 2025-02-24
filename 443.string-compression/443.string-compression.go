/*
 * @lc app=leetcode.cn id=443 lang=golang
 *
 * [443] String Compression
 */

// @lc code=start
func compress(chars []byte) int {
	nums := []byte("0123456789")
	count := 1
	pre := 0
	for i := 1; i < len(chars)+1; i++ {
		if i == len(chars) || chars[pre] != chars[i] {
			if count == 1 {
				pre = pre + 1
			} else if count <= 9 {
				chars[pre+1] = nums[count]
				pre = pre + 2
			} else if count <= 99 {
				chars[pre+1] = nums[count/10]
				chars[pre+2] = nums[count%10]
				pre = pre + 3
			} else if count <= 999 {
				chars[pre+1] = nums[count/100]
				chars[pre+2] = nums[(count%100)/10]
				chars[pre+3] = nums[count%10]
				pre = pre + 4
			} else {
				chars[pre+1] = nums[count/1000]
				chars[pre+2] = nums[(count%1000)/100]
				chars[pre+3] = nums[(count%100)/10]
				chars[pre+4] = nums[count%10]
				pre = pre + 5
			}
			if i < len(chars) {
				chars[pre] = chars[i]
				count = 1
			}
		} else if i < len(chars) && chars[pre] == chars[i] {
			count++
		}
	}
	// fmt.Println(chars[:pre])
	return pre
}

// @lc code=end
