package main

import (
	"fmt"
	"strconv"
)

func findDigits(n int32) int32 {
    // Write your code here
	var res int32

	s := strconv.Itoa(int(n))

	runes := []rune(s)

	digits := make([]int, len(runes))

	
	for i, r := range runes {
		digits[i] = int(r - '0')
	}
	
	for i := 0; i < len(digits); i++ {
		if digits[i] != 0 && n % int32(digits[i]) == 0 {
			res++
		}
	}

	return res
}

func main() {
	fmt.Println(findDigits(124)) // 3
	fmt.Println(findDigits(111)) // 3
	fmt.Println(findDigits(10)) // 1
	fmt.Println(findDigits(12)) // 2
	fmt.Println(findDigits(1012)) // 3
	fmt.Println(findDigits(123456789)) // 3
	fmt.Println(findDigits(114108089)) // 6
}