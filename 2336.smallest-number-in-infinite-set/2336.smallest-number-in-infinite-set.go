/*
 * @lc app=leetcode.cn id=2336 lang=golang
 *
 * [2336] Smallest Number in Infinite Set
 */

// @lc code=start
type SmallestInfiniteSet struct {
	Heap []int
	Set  map[int]struct{}
}

func Constructor() SmallestInfiniteSet {
	set := SmallestInfiniteSet{Heap: make([]int, 0), Set: make(map[int]struct{})}
	for i := 1; i <= 1000; i++ {
		set.AddBack(i)
	}
	return set
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	// get smallest and swap latest element
	smallest := this.Heap[0]
	this.Heap[0] = this.Heap[this.Length()-1]

	//delete from set and heap
	delete(this.Set, smallest)
	this.Heap = this.Heap[:this.Length()-1]

	this.siftDown(0)
	return smallest
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	_, exists := this.Set[num]
	if exists {
		return
	}
	this.Set[num] = struct{}{}
	this.Heap = append(this.Heap, num)
	this.siftUp(this.Length() - 1)
}

func (this *SmallestInfiniteSet) Parent(index int) int {
	index = (index - 1) / 2
	if index >= 0 {
		return index
	}
	return -1
}

func (this *SmallestInfiniteSet) Left(index int) int {
	index = 2*index + 1
	if index > this.Length()-1 {
		return -1
	}
	return index
}

func (this *SmallestInfiniteSet) Right(index int) int {
	index = 2*index + 2
	if index > this.Length()-1 {
		return -1
	}
	return index
}

func (this *SmallestInfiniteSet) Length() int {
	return len(this.Heap)
}

func (this *SmallestInfiniteSet) siftDown(index int) {
	for {
		if this.Right(index) != -1 && this.Heap[index] > this.Heap[this.Right(index)] && this.Heap[this.Left(index)] > this.Heap[this.Right(index)] {
			this.Heap[index], this.Heap[this.Right(index)] = this.Heap[this.Right(index)], this.Heap[index]
			index = this.Right(index)
		} else if this.Left(index) != -1 && this.Heap[index] > this.Heap[this.Left(index)] {
			this.Heap[index], this.Heap[this.Left(index)] = this.Heap[this.Left(index)], this.Heap[index]
			index = this.Left(index)
		} else {
			break
		}
	}
}

func (this *SmallestInfiniteSet) siftUp(index int) {
	for index > 0 && this.Heap[this.Parent(index)] > this.Heap[index] {
		this.Heap[this.Parent(index)], this.Heap[index] = this.Heap[index], this.Heap[this.Parent(index)]
		index = this.Parent(index)
	}
}

/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */
/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */
// @lc code=end
