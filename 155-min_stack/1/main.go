package main

import (
	"fmt"
)

func main() {
	stack := Constructor()
	stack.Push(512)
	stack.Push(-1024)
	stack.Push(-1024)
	stack.Push(512)
	stack.Pop()
	fmt.Println(stack.GetMin())
	stack.Pop()
	fmt.Println(stack.GetMin())
	stack.Pop()
	printStacks(stack)
	fmt.Println(stack.GetMin())
	printStacks(stack)
}

func printStacks(s MinStack) {
	fmt.Printf("stack: %v - min stack: %v\n", s.stack, s.minStack)
}

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if len(this.minStack) == 0 {
		this.minStack = append(this.minStack, val)
	} else {
		this.minStack = append(this.minStack, min(this.minStack[len(this.minStack)-1], val))
	}
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}
