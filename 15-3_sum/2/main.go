package main

import "sort"

func main() {

}

func threeSum(nums []int) [][]int {
	res := [][]int{}

	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if i == 0 || nums[i] != nums[i-1] {
			// 2 sum
			low, high := i+1, len(nums)-1
			sum := 0 - nums[i] // sum

			for low < high {
				if nums[low]+nums[high] == sum { // found 3 sum
					res = append(res, []int{nums[i], nums[low], nums[high]})

					// encontrando elementos unicos low
					for low < high && nums[low] == nums[low+1] {
						low++
					}

					// encontrando os elementos unicos high
					for high > low && nums[high] == nums[high-1] {
						high--
					}

					high--
					low++

					// se forem maiores que a sum, diminuimos o high, como o array está ordenado, diminuirá a soma
				} else if nums[low]+nums[high] > sum {
					high--
				} else {
					low++
				}
			}
		}
	}

	return res
}
