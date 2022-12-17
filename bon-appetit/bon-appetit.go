package main

import "fmt"

func bonAppetit(bill []int32, k int32, b int32) {
	// Write your code here
	var sum int32

	for i, v := range bill {
		if int32(i) != k {
			sum += v
		}
	}

	sum = sum / 2

	if b == sum {
		fmt.Println("Bon Appetit")
	} else {
		fmt.Println(b - sum)
	}
}

func main() {
	bonAppetit([]int32{3, 10, 2, 9}, 1, 7) 		// Bon Appetit
	bonAppetit([]int32{3, 10, 2, 9}, 1, 12) 	// 5
}