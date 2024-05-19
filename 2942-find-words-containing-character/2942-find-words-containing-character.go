func findWordsContaining(words []string, x byte) []int {
	var result []int

	for i := 0; i < len(words); i++ {
		for _, chars := range words[i] {
			if string(x) == string(chars) {
				result = append(result, i)
				break
			}
		}
	}

	return result
}