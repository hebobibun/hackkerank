package main

import "fmt"

func chocolateFeast(n int32, c int32, m int32) int32 {
    // Write your code here
	var sum int32

	sum = n / c

	if sum >= m {
		sisa := (sum % m) + (sum / m)
		sum = sum + (sum / m)
		if sisa >= m {
			sum = sum + (sisa / m)
		}
	}

	return sum
}

func main() {
	fmt.Println(chocolateFeast(10, 2, 5)) // 6
	fmt.Println(chocolateFeast(12, 4, 4)) // 3
	fmt.Println(chocolateFeast(6, 2, 2)) // 5
}