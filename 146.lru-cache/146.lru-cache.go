/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU Cache
 */
package main

// @lc code=start
type Node struct {
	Pre   *Node
	Next  *Node
	Key   int
	Value int
}

type LRUCache struct {
	Size       int
	Capacity   int
	cache      map[int]*Node
	head, tail *Node // dummy node
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.Next = tail
	tail.Pre = head
	return LRUCache{Size: 0, Capacity: capacity, cache: map[int]*Node{}, head: head, tail: tail}
}

func (this *LRUCache) Get(key int) int {
	_, exists := this.cache[key]
	if !exists {
		return -1
	}
	node := this.cache[key]
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre

	node.Next = this.head.Next
	this.head.Next.Pre = node
	node.Pre = this.head
	this.head.Next = node

	return node.Value
}

func (this *LRUCache) Put(key int, value int) {
	node, exists := this.cache[key]

	if exists {
		node.Value = value
		this.Get(key)
		return
	}

	if len(this.cache) == this.Capacity {
		node = this.tail.Pre
		delete(this.cache, node.Key)

		pre := node.Pre
		next := node.Next

		pre.Next = next
		next.Pre = pre

		node.Key = key
		node.Value = value
	} else {
		node = &Node{Key: key, Value: value}
	}
	this.cache[key] = node

	node.Next = this.head.Next
	this.head.Next.Pre = node
	node.Pre = this.head
	this.head.Next = node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
