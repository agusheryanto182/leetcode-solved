func singleNumber(nums []int) int {
	for i, numb := range nums {
		for j := 0; j < len(nums); j++ {
			if numb == nums[j] && i != j {
				break
			}

			if j == len(nums)-1 {
				return numb
			}
		}
	}
	return 0
}
