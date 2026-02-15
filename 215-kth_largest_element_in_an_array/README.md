Esse exercício podemos usar **QuickSelect e MinHeap**.

https://medium.programmerscareer.com/leetcode-215-golang-kth-largest-element-in-an-array-insights-into-quickselect-and-minheap-5342963d1505

## MinHeap

O quickselect é uma variação do QuickSort, e apesar de ser mais eficiente, vou fazer com MinHeap.

Basicamente, vamos usar uma **minHeap**. Nesse caso:

1. Inicializamo um min heap com os nums pré populados
2. fazemos pop até len min heap < k
3. retornamos o pop

```go
func findKthLargest(nums []int, k int) int {
	hp := minHeap(nums)
	heap.Init(&hp)
	for hp.Len() > k {
		heap.Pop(&hp)
	}
	return heap.Pop(&hp).(int)
}
```

Nossa heap, ficará normal:

```go
minHeap []int

func (m *minHeap) Push(x any) {
	*m = append(*m, x.(int))
}

func (m *minHeap) Pop() any {
	old := *m
	size := len(*m)
	x := old[size-1]
	*m = old[:size-1]
	return x
}

func (m *minHeap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *minHeap) Len() int {
	return len(*m)
}

func (m *minHeap) Less(i, j int) bool {
	return (*m)[i] < (*m)[j]
}
```