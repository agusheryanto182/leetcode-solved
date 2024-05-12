func numIdenticalPairs(nums []int) int {
	counter := 0
	temp := make(map[int]int)
	for t, numb := range nums {
		temp[t] = numb
	}


	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == temp[j] {
				counter++
			} else if nums[i] == nums[j] {
				counter++
			}
		}
	}
	return counter
}