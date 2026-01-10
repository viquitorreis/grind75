package main

import (
	"fmt"
)

func main() {
	fmt.Println("bad version: ", firstBadVersion(8))
}

func isBadVersion(v int) bool {
	versions := map[int]bool{
		1: false,
		2: false,
		3: false,
		4: false,
		5: false,
		6: true,
		7: true,
		8: true,
	}

	val, ok := versions[v]
	fmt.Println("val is: ", val)
	if ok {
		return val
	}

	return false
}

func firstBadVersion(n int) int {
	if n == 1 {
		return n
	}

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
