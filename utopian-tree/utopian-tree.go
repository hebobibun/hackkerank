package main

import "fmt"

func utopianTree(n int32) int32 {
	// Write your code here
	res := 1

	for i := 1; i <= int(n); i++ {
		if i % 2 == 0 {
			res += 1
		} else if i % 2 == 1 {
			res *= 2
		}
	}

	return int32(res)
}

func main() {
	fmt.Println(utopianTree(5)) // 14
	fmt.Println(utopianTree(0)) // 1
	fmt.Println(utopianTree(1)) // 2
	fmt.Println(utopianTree(2)) // 3
	fmt.Println(utopianTree(4)) // 7
}