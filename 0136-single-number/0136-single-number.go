func singleNumber(nums []int) int {
	temp := make(map[int]int)
	for i, numb := range nums {
		for j := 0; j < len(nums); j++ {
			if numb == nums[j] && i != j {
				temp[numb] = i
			}
		}
	}

	for _, char := range nums {
		_, ok := temp[char]
		if !ok {
			return char
		}
	}
	return 0
}