package main

import "fmt"

func main() {
	fmt.Println(firstBadVersion(7))
}

func isBadVersion(version int) bool {
	versions := map[int]bool{
		1: false,
		2: false,
		3: false,
		4: false,
		5: true,
		6: true,
		7: true,
	}

	if val, ok := versions[version]; ok {
		return val
	}

	panic("out of bounds version")
}

func firstBadVersion(n int) int {
	left, right := 1, n

	for left < right {
		mid := (left + right) >> 1

		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}
