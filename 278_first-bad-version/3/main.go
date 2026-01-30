package main

import (
	"fmt"
)

func main() {
	fmt.Println(firstBadVersion(4))
}

func isBadVersion(version int) bool {
	fmt.Println("checking: ", version)

	vs := map[int]bool{
		1: true,
		2: true,
		3: true,
		4: true,
		5: true,
	}

	return vs[version]
}

func firstBadVersion(n int) int {
	// se é bad, buscamos a esquerda -> pode ter uma anterior bad também
	// se ainda não é bad, buscamos a direita
	// versions := make(map[int]bool, n)
	// seen := make(map[int]bool, n)
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
