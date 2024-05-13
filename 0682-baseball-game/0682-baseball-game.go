func calPoints(operations []string) int {
	record := []int{}

	for _, o := range operations {
		if o == "C" {
			record = record[:len(record)-1]
		} else if o == "D" {
			record = append(record, 2*record[len(record)-1])
		} else if o == "+" {
			record = append(record, record[len(record)-2]+record[len(record)-1])
		} else {
			numb, _ := strconv.Atoi(o)
			record = append(record, numb)
		}
	}

	res := 0
	for _, numb := range record {
		res += numb
	}
	return res
}