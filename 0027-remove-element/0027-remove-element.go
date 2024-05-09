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
