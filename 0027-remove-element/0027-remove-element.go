package main

import "fmt"

func removeElement(nums []int, val int) int {
	k := 0

	for _, numb := range nums {
		if numb != val {
			nums[k] = numb
			k++
		}
	}

	fmt.Println(nums)
	return k
}

func main() {
	input1 := removeElement([]int{3, 2, 2, 3}, 3)
	fmt.Println(input1)
}
