func scoreOfString(s string) int {
	count := 0

	for i := 0; i < len(s)-1; i++ {
		curr := int(math.Abs(float64((int(s[i]) - int(s[i+1])))))
		count += curr
	}
	return count
}