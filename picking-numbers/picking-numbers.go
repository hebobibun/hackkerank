package main

import (
	"fmt"
	"sort"
)

func pickingNumbers(a []int32) int32 {
    // Write your code here
	var res int32

	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})

	slc := []int32{a[0]}
	countSLC := []int32{}

	fmt.Println(a)

	for i := 0; i < len(a); i++ {
		if i == 0 {
			continue
		} else if a[i] == slc[0] || a[i] == slc[0] + 1 {
			slc = append(slc, a[i])
			fmt.Println("slc after append : ", slc)
		} else if a[i] != slc[0] || a[i] != slc[0] + 1 {
			fmt.Println("slc before count : ", slc)
			countSLC = append(countSLC, int32(len(slc)))
			slc = []int32{a[i]}
		}

		countSLC = append(countSLC, int32(len(slc)))
	}

	for _, v := range countSLC {
		if v > res {
			res = v
		}
	}

	// fmt.Println(a)
	// fmt.Println(countSLC)

	return res
}

func main() {
	fmt.Println(pickingNumbers([]int32{4, 6, 5, 3, 3, 1})) // 3
	fmt.Println(pickingNumbers([]int32{5, 6, 6, 6, 3, 1})) // 4
	fmt.Println(pickingNumbers([]int32{1, 2, 2, 3, 1, 2})) // 5
}