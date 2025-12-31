package main

import (
	"fmt"
)

func main() {
	q := Constructor()
	fmt.Printf("peek: %d\n", q.Peek())
	q.Push(1)
	q.Push(2)
	fmt.Printf("peek: %d\n", q.Peek())
	param1 := q.Pop()
	fmt.Printf("param pop: %d\n", param1)
}

type MyQueue struct {
	q   []int
	len int
}

func Constructor() MyQueue {
	return MyQueue{
		q:   make([]int, 0),
		len: 0,
	}
}

// push -> append
func (this *MyQueue) Push(x int) {
	this.q = append(this.q, x)
	this.len++
}

// pop: must remove FRONT(oldest) element and return this element
func (this *MyQueue) Pop() int {
	if this.Empty() {
		return -1
	}

	el := this.q[0]
	this.q = this.q[1:this.len]
	this.len--

	return el
}

// peek -> must see first element from q
func (this *MyQueue) Peek() int {
	if this.len == 0 {
		return 0
	}

	return this.q[0]
}

func (this *MyQueue) Empty() bool {
	return this.len == 0
}
