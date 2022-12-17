package main

import (
	"fmt"
	"sort"
)

func sockMerchant(n int32, ar []int32) int32 {
	// Write your code here
	var count int32

	sort.Slice(ar, func(i, j int) bool {
		return ar[i] < ar[j]
	})

	fmt.Println(ar)

	pairs := []int32{ar[0]}

	for i, v := range ar {
		if i == 0 {
			continue
		} else if len(pairs) == 0 {
			pairs = append(pairs, v)
		} else if v == pairs[0] {
			pairs = append(pairs, v)
		} else {
			pairs = []int32{ar[i]}
		}

		fmt.Println(i, pairs)

		if len(pairs) == 2 {
			pairs = []int32{}
			count += 1
		}
	}

	return count
}

func main() {
	fmt.Println(sockMerchant(9, []int32{10, 20, 20, 10, 10, 30, 50, 10, 20})) // 3
									// [10 10 10 10 20 20 20 30 50]
	fmt.Println(sockMerchant(10, []int32{1, 1, 3, 1, 2, 1, 3, 3, 3, 3})) // 4
									// [1 1 1 1 2 3 3 3 3 3]
	fmt.Println(sockMerchant(10, []int32{1, 4, 3, 4, 2, 2, 4, 4, 3, 3, 3})) // 5
									// [1 1 1 1 2 3 3 3 3 3]
}