package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var dict = make(map[int]int)
	for indexCurr, valCurr := range nums {
		indexDict, isPresent := dict[target-valCurr]

		if isPresent {
			return []int{indexDict, indexCurr}
		}
		dict[valCurr] = indexCurr
	}

	return []int{}
}

func main() {
	input1 := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(input1)

	input2 := twoSum([]int{3, 2, 4}, 6)
	fmt.Println(input2)

	input3 := twoSum([]int{3, 3}, 6)
	fmt.Println(input3)
}
