func numIdenticalPairs(nums []int) int {
	counter := 0
	temp := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			numb, ok := temp[j]
			if ok && numb == nums[i] {
				counter++
			} else if nums[i] == nums[j] {
				counter++
			}
			temp[j] = nums[j]
		}
	}
	return counter
}