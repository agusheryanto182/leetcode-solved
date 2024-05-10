func missingNumber(nums []int) int {
	if len(nums) == 1 {
		if nums[0] != 0 {
			return nums[0] - 1
		}
		return nums[0] + 1
	}

	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if nums[i] > nums[j] {
				temp := nums[i]
				nums[i] = nums[j]
				nums[j] = temp
			}
		}
	}

	temp := 0
	for k := 0; k < len(nums)-1; k++ {
		if nums[0] != 0 {
			return nums[0] - 1
		}
		if nums[k]+1 != nums[k+1] {
			return nums[k] + 1
		}
		temp = nums[k+1] + 1
	}
	return temp
}
