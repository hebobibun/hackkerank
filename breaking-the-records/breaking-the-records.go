package main

import "fmt"

func breakingRecords(scores []int32) []int32 {
    // Write your code here
	var res []int32
	low := scores[0]
	high := scores[0]

	var clow int32
	var chigh int32

	for i := 0; i < len(scores); i++ {
		if scores[i] > high {
			high = scores[i]
			chigh++
		} else if scores[i] < low {
			low = scores[i]
			clow++
		}
	}

	res = append(res, chigh)
	res = append(res, clow)

	return res
}

func main() {
	fmt.Println(breakingRecords([]int32{10, 5, 20, 20, 4, 5, 2, 25, 1})) // [2, 4]
	fmt.Println(breakingRecords([]int32{3, 4, 21, 36, 10, 28, 35, 5, 24, 42})) // [4, 0]
}