func scoreOfString(s string) int {
	count := 0
	asciStr := []int{}
	for _, char := range s {
		asciStr = append(asciStr, int(char))
	}

	for i := 0; i < len(s)-1; i++ {
		curr := int(math.Abs(float64((asciStr[i] - asciStr[i+1]))))
		count += curr
	}
	return count
}