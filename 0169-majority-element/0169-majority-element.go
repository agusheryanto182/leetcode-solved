func majorityElement(nums []int) int {
	dict := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		dict[nums[i]]++
	}

	result := 0
	temp := 0
	for _, numb := range nums {
		if dict[numb] > temp {
			temp = dict[numb]
			result = numb
		}
	}

	return result
}