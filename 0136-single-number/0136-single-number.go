func singleNumber(nums []int) int {
	occurrence := make(map[int]int)

	// Menghitung kemunculan setiap angka
	for _, num := range nums {
		occurrence[num]++
	}

	// Mencari angka tunggal
	for num, count := range occurrence {
		if count == 1 {
			return num
		}
	}

	return 0 // Jika tidak ada angka tunggal, maka mengembalikan 0
}