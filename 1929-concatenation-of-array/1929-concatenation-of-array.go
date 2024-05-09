func getConcatenation(nums []int) []int {
	// create new array
	var newArray []int

	// do looping
	i := 0
	for i != 2 {
		newArray = append(newArray, nums...)
		i++
	}
	return newArray
}