/*
 * @lc app=leetcode.cn id=232 lang=golang
 *
 * [232] Implement Queue using Stacks
 */

// @lc code=start
type MyQueue struct {
	pushStack []int
	popStack  []int
}

func Constructor() MyQueue {
	return MyQueue{pushStack: []int{}, popStack: []int{}}
}

func (this *MyQueue) transfer() {
	// transfer from push stack to pop stack
	if len(this.popStack) > 0 {
		return
	}

	for len(this.pushStack) > 0 {
		// pop stack
		top := this.pushStack[len(this.pushStack)-1]
		this.pushStack = this.pushStack[:len(this.pushStack)-1]
		this.popStack = append(this.popStack, top)
	}
}

func (this *MyQueue) Push(x int) {
	this.pushStack = append(this.pushStack, x)
}

func (this *MyQueue) Pop() int {
	if len(this.popStack) == 0 {
		this.transfer()
	}

	top := this.popStack[len(this.popStack)-1]
	this.popStack = this.popStack[:len(this.popStack)-1]
	return top
}

func (this *MyQueue) Peek() int {
	if len(this.popStack) == 0 {
		this.transfer()
	}
	return this.popStack[len(this.popStack)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.pushStack) == 0 && len(this.popStack) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
// @lc code=end

