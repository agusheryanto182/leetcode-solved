
func plusOne(digits []int) []int {
	// Iterate from the least significant digit to the most significant digit
	for i := len(digits) - 1; i >= 0; i-- {
		// Increment the current digit by 1
		digits[i]++

		// If the digit is less than or equal to 9, no carry-over, return digits
		if digits[i] <= 9 {
			return digits
		}

		// Otherwise, set the current digit to 0 and continue to carry-over to the next digit
		digits[i] = 0
	}

	// If we reach here, it means there's a carry-over to the most significant digit
	// We need to insert a new digit 1 at the beginning of the slice
	return append([]int{1}, digits...)
}
