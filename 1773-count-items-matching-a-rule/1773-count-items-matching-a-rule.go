func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	result := 0
	rule := make(map[string]int)
	rule["type"] = 0
	rule["color"] = 1
	rule["name"] = 2

	for i := range items {
		if items[i][rule[ruleKey]] == ruleValue {
			result++
		}
	}
	return result
}