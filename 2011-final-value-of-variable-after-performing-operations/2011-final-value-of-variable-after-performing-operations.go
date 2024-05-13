func finalValueAfterOperations(operations []string) int {
	result := 0
	for _, char := range operations {
		if char == "X++" || char == "++X" {
			result = result + 1
		} else {
			result = result - 1
		}
	}
	return result
}
