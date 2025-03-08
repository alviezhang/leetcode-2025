/*
 * @lc app=leetcode.cn id=901 lang=golang
 *
 * [901] Online Stock Span
 */
// @lc code=start
type StockInfo struct {
	Price int
	Index int
}

type StockSpanner struct {
	count int
	stack []StockInfo
}

func Constructor() StockSpanner {
	return StockSpanner{stack: []StockInfo{}}
}

func (this *StockSpanner) Next(price int) int {
	var ans int
	for len(this.stack) > 0 && this.stack[len(this.stack)-1].Price <= price {
		this.stack = this.stack[:len(this.stack)-1]
	}

	if len(this.stack) == 0 {
		ans = this.count + 1
	} else {
		ans = this.count - this.stack[len(this.stack)-1].Index
	}

	this.stack = append(this.stack, StockInfo{Price: price, Index: this.count})
	this.count++
	return ans
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
// @lc code=end

