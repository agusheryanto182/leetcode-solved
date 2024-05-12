func numIdenticalPairs(nums []int) int {
	counter := 0
	temp := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		temp[nums[i]]++
		counter += temp[nums[i]] - 1
	}
	return counter
}