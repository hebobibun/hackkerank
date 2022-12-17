package main

import "fmt"

func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
    // Write your code here
	var sum int32

	for i := int32(0); i < int32(n); i++ {
		for j := int32(0); j < int32(n); j++ {
			if i < j && (ar[i] + ar[j]) % k == 0 {
				fmt.Println(ar[i], "+", ar[j], "=", ar[i] + ar[j])
				sum++
			}
		}
	}
	return sum
}

func main() {
	fmt.Println(divisibleSumPairs(6, 4, []int32{1, 3, 2, 6, 1, 2})) // 5
}