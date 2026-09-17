package main

import (
	"fmt"
)

func romanToInt(s string) int {
	roman := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0

	for i := 0; i < len(s); i++ {
		current := roman[s[i]]
		if i != len(s)-1 && current < roman[s[i+1]] {
			result -= current
		} else {
			result += current
		}
	}

	return result
}

func main() {
	input1 := romanToInt("III")
	fmt.Println(input1)

	input2 := romanToInt("LVIII")
	fmt.Println(input2)

	input3 := romanToInt("MCMXCIV")
	fmt.Println(input3)
}
