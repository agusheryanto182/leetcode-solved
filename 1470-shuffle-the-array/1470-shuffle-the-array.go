func shuffle(nums []int, n int) []int {
	result := []int{}
	numsx := []int{}
	numsy := []int{}
	numsx = append(numsx, nums[:len(nums)/2]...)
	numsy = append(numsy, nums[len(nums)/2:]...)

	for i := 0; i < len(nums)/2; i++ {
		result = append(result, numsx[i])
		result = append(result, numsy[i])
	}
	return result
}