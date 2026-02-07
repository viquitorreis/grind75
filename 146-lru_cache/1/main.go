package main

func main() {

}

// head is the LRU
type Node struct {
	Key, Val   int
	Prev, Next *Node
}

type LRUCache struct {
	cache map[int]*Node
	// head e tail são dummy nodes, eles nao armazenam dados reais, sao sentinelas
	head, tail *Node
	capacity   int
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.Next = tail
	tail.Prev = head

	return LRUCache{
		cache:    make(map[int]*Node),
		capacity: capacity,
		head:     head,
		tail:     tail,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.Remove(node)
		this.AddToFront(node)
		return node.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		// so atualiza o valor
		node.Val = value
		this.Remove(node)
		this.AddToFront(node)
		return
	}

	if len(this.cache) == this.capacity {
		lru := this.tail.Prev
		this.Remove(lru)
		delete(this.cache, lru.Key)
	}

	newNode := &Node{
		Key: key,
		Val: value,
	}
	this.cache[key] = newNode
	this.AddToFront(newNode)
}
func (this *LRUCache) Remove(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (this *LRUCache) AddToFront(node *Node) {
	node.Next = this.head.Next
	node.Prev = this.head
	this.head.Next.Prev = node
	this.head.Next = node
}
