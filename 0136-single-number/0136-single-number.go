
func singleNumber(nums []int) int {
	// 1, 2, 3, 2, 1
	// 3:  0 1 1
	// 1:  0 0 1
	// ------
	// XOR 0 1 0

	var result int

	for _, n := range nums {
		result = result ^ n
	}

	return result
}