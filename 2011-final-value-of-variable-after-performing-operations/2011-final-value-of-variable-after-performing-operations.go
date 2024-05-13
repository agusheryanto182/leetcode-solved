func finalValueAfterOperations(operations []string) int {
	result := 0
	for _, char := range operations {
		if char[1] == '+' {
			result++
		} else {
			result--
		}
	}
	return result
}