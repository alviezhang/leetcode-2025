/*
 * @lc app=leetcode.cn id=933 lang=golang
 *
 * [933] Number of Recent Calls
 */

// @lc code=start
type RecentCounter struct {
	pingQueue []int
}

func Constructor() RecentCounter {
	return RecentCounter{pingQueue: make([]int, 0)}
}

func (this *RecentCounter) Ping(t int) int {
	this.pingQueue = append(this.pingQueue, t)
	for len(this.pingQueue) != 1 && t-this.pingQueue[0] > 3000 {
		this.pingQueue = this.pingQueue[1:]
	}
	return len(this.pingQueue)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
// @lc code=end

