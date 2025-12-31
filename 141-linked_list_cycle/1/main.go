package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// node4 := &ListNode{Val: -4}
	// node3 := &ListNode{Val: 0}
	// node2 := &ListNode{Val: 2}
	// node1 := &ListNode{Val: 3}

	// node1.Next = node2
	// node2.Next = node3
	// node3.Next = node4
	// node4.Next = node2

	// println("has cycle 1? ", hasCycle(node1))

	// node5 := &ListNode{Val: 1}
	// node6 := &ListNode{Val: 2}

	// node5.Next = node6
	// node6.Next = node5

	// println("has cycle 2? ", hasCycle(node5))

	// node7 := &ListNode{Val: 1}
	// node7.Next = &ListNode{
	// 	Val:  -1,
	// 	Next: nil,
	// }

	// println("has cycle 3? ", hasCycle(node7))

	// node0 := &ListNode{Val: -1}
	// node1 := &ListNode{Val: -7}
	// node2 := &ListNode{Val: 7}
	// node3 := &ListNode{Val: -4}
	// node4 := &ListNode{Val: 19}
	// node5 := &ListNode{Val: 6}
	// node6 := &ListNode{Val: -9}
	// node7 := &ListNode{Val: -5}
	// node8 := &ListNode{Val: -2}
	// node9 := &ListNode{Val: -5}

	// node0.Next = node1
	// node1.Next = node2
	// node2.Next = node3
	// node3.Next = node4
	// node4.Next = node5
	// node5.Next = node6
	// node6.Next = node7
	// node7.Next = node8
	// node8.Next = node9
	// node9.Next = node6

	// println("has cycle?", hasCycle(node0))

	node0 := &ListNode{Val: -21}
	node1 := &ListNode{Val: 10}
	node2 := &ListNode{Val: 17}
	node3 := &ListNode{Val: 8}
	node4 := &ListNode{Val: 4}
	node5 := &ListNode{Val: 26}
	node6 := &ListNode{Val: 5}
	node7 := &ListNode{Val: 35}
	node8 := &ListNode{Val: 33}
	node9 := &ListNode{Val: -7}
	node10 := &ListNode{Val: -16}
	node11 := &ListNode{Val: 27}
	node12 := &ListNode{Val: -12}
	node13 := &ListNode{Val: 6}
	node14 := &ListNode{Val: 29}
	node15 := &ListNode{Val: -12}
	node16 := &ListNode{Val: 5}
	node17 := &ListNode{Val: 9}
	node18 := &ListNode{Val: 20}
	node19 := &ListNode{Val: 14}
	node20 := &ListNode{Val: 14}
	node21 := &ListNode{Val: 2}
	node22 := &ListNode{Val: 13}
	node23 := &ListNode{Val: -24}
	node24 := &ListNode{Val: 21}
	node25 := &ListNode{Val: 23}
	node26 := &ListNode{Val: -21}
	node27 := &ListNode{Val: 5}

	node0.Next = node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	node6.Next = node7
	node7.Next = node8
	node8.Next = node9
	node9.Next = node10
	node10.Next = node11
	node11.Next = node12
	node12.Next = node13
	node13.Next = node14
	node14.Next = node15
	node15.Next = node16
	node16.Next = node17
	node17.Next = node18
	node18.Next = node19
	node19.Next = node20
	node20.Next = node21
	node21.Next = node22
	node22.Next = node23
	node23.Next = node24
	node24.Next = node25
	node25.Next = node26
	node26.Next = node27
	node27.Next = node24 // pos = 24, cycle back to node24 (21)

	println("has cycle?", hasCycle(node0))
}

// existe um ciclo na linked list. Onde um node da lista pode ser encontrado novamente
// ao seguir o next indefinidamente
// pos -> variavel global que fala para o indice do node que a TAIL aponta
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	// listNode memory address -> count
	seen := make(map[string]int, 0)

	return travess(head, seen)
}

var count = 0

func travess(curr *ListNode, prevSeen map[string]int) bool {
	if curr.Next == nil {
		return false
	}

	if curr != nil {
		count++
		currAddr := fmt.Sprintf("%s", curr)
		nextAddr := fmt.Sprintf("%s", curr.Next)

		if prevSeen[currAddr] != 0 {
			fmt.Printf("current val: %d - count: %d \n", curr.Val, prevSeen[currAddr])
			// fora da lista -> nao ciclica
			if prevSeen[nextAddr] == 0 {
				fmt.Printf("Next.Val: %d - count: %d\n", curr.Next.Val, prevSeen[nextAddr])
				return false
			}
		}

		prevSeen[currAddr]++

		if prevSeen[currAddr] == 2 {
			// has cycle :)
			if prevSeen[nextAddr] >= 2 {
				return true
			}
		}

		fmt.Printf("travess: curr %d - count: %d - prevSeen: %v\n", curr.Val, count, prevSeen)

	}

	curr = curr.Next

	return travess(curr, prevSeen)
}
