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

// Complexity Analysis
// BIG O
//		- O(n)
//			- Push: O(1) - in Go most of the times (unless underlying array needs to resize)
//			- Pop: O(1) - the operation this.q = this.q[1:] creates a new SLICE HEADER pointing to the same underlying array
// 				starting from index 1, that is constant time. However, there is a memory cost here, the old element isnt freed
//				from the underlying array until the whole array is garbage collected, which can lead to memory waste over time
//			- Peek: O(1)
//			- Empty: O(1) - simple comparisons

// SPACE COMPLEXITY
//		- O(n) - where n is the elements currently stored in the queue

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

// Pop -> Removes the element from the front of the queue and returns it.
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

// Peek -> returns the element at the front of the queue.
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
