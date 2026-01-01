package main

import "fmt"

func main() {
	q := Constructor()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)
	q.Push(5)
	printQueue(&q)
	q.Pop()
	q.Pop()
	q.Pop()
	q.Pop()
	q.Pop()
}

/**
* Your MyQueue object will be instantiated and called as such:
* obj := Constructor();
* obj.Push(x);
* param_2 := obj.Pop();
* param_3 := obj.Peek();
* param_4 := obj.Empty();
 */

type MyQueue struct {
	q   []int
	len int
}

func Constructor() MyQueue {
	return MyQueue{
		q:   []int{},
		len: 0,
	}
}

func (this *MyQueue) Push(x int) {
	this.q = append(this.q, x)
	this.len++
}

// Removes the element from the front of the queue and returns it.
func (this *MyQueue) Pop() int {
	if this.Empty() {
		return -1
	}

	el := this.q[0]

	if this.len == 1 {
		this.q = []int{}
	} else {
		this.q = this.q[1:]
	}

	this.len--
	return el
}

// returns the element at the front of the queue.
func (this *MyQueue) Peek() int {
	if this.Empty() {
		return -1
	}

	return this.q[0]
}

func (this *MyQueue) Empty() bool {
	return this.len == 0
}

func printQueue(q *MyQueue) {
	for i := range q.len {
		fmt.Printf("%d ", q.q[i])
	}
	println()
}
