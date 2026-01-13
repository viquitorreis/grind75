package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	list2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	mergeTwoLists(list1, list2)
}

// ao mergear duas listas, precisamos:
// 1. monitorar quem está a frente (next) - head
// 2. monitorar quem é o último elemento atual da lista, para sabermos onde colocar a próxima pessoa - tail
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	tail := head // inicialmente o mesmo endereço da head. Essencial para alterarmos o ponteiro do head.Next depois

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			tail.Next = list1  // na primeira iteração vai alterar o Next da head também
			list1 = list1.Next // movendo a linked list da lista 1
		} else {
			tail.Next = list2  // na primeira iteração vai alterar o Next da head também
			list2 = list2.Next // movendo a linked list da lista 2
		}

		// aqui a tail já vai ficar com ponteiro diferente da head
		tail = tail.Next
	}

	// ao terminar, no maximo uma lista vai ter mais um node depois de quebrar a condição do loop
	if list1 != nil {
		tail.Next = list1
	}
	if list2 != nil {
		tail.Next = list2
	}

	printList(head.Next)

	return head.Next
}

func printList(list *ListNode) {
	for list != nil {
		fmt.Printf("val: %d", list.Val)
		if list.Next != nil {
			fmt.Printf(" next: %d ", list.Next.Val)
		}

		list = list.Next
	}

	println()
}
