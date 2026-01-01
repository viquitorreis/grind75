package main

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	println(majorityElement(nums))
}

// solução interessante: https://leetcode.com/problems/majority-element/solutions/3220196/beats-9987-with-go-by-ratchawat-1rtu/

// Complexity Analisys
// BIG O - O(n) - linear time
// 		loop on array only once, and most other operations (variables, conditionals, etc) are O(1) - constant
// Space Complexity O(n) - linear
//		Worst case scenario: all elements on array are different, so my hashmap will have one different key for each, so n entries on map
//		Of course, asymptotically

func majorityElement(nums []int) int {
	seen := make(map[int]int, 0)
	var (
		biggestNum   = 0
		biggestOccur = 0
	)

	for _, num := range nums {
		seen[num]++
		if count, ok := seen[num]; ok {
			if biggestOccur < count {
				biggestOccur = seen[num]
				biggestNum = num
			}
		}
	}

	return biggestNum
}
