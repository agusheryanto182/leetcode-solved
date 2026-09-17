package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	reversedInt := reversedInt(x)

	if x != reversedInt {
		return false
	}
	return true
}

func reversedInt(n int) int {
	result := 0

	for n != 0 {
		digit := n % 10
		result = result*10 + digit
		n /= 10
	}

	return result
}

func main() {
	input1 := isPalindrome(121)
	fmt.Println(input1)

	input2 := isPalindrome(-121)
	fmt.Println(input2)
}
