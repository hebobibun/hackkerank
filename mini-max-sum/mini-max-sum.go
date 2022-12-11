package main

import (
	"fmt"
)

/*
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func miniMaxSum(arr []int32) {
	var numMin, numMax, val, numSum int32

	for i, v := range arr {
		val = v
		if numMax < val {
			numMax = val
		}

		if i == 0 {
			numMin = val
		} else {
			if numMin > val {
				numMin = val 
			}
		}
		numSum += val
	}

	fmt.Println(numSum - numMax, numSum - numMin)
}

func main() {
	miniMaxSum([]int32{1, 2, 3, 4, 5}) // 10, 14
	miniMaxSum([]int32{7, 69, 2, 221, 8974}) // 299, 9271
}

	
	
	// func miniMaxSum(arr []int32) {
	//     // Write your code here
	// 	var slc = []int32{}
	// 	var tmp int32
	
	// 	sort.Slice(arr, func(i, j int) bool {
	// 		return arr[i] < arr[j]
	// 	})
	
	// 	for j := 0; j < len(arr); j++ {
	// 		for k := 0; k < len(arr); k++ {
	// 			if k == j {
	// 				continue
	// 			} else {
	// 				tmp += arr[k]
	// 			}
	// 		}
	// 		slc = append(slc, tmp)
	// 		tmp = 0
	// 	}
	
	// 	min := slc[0]
	// 	max := slc[0]
	
	// 	for _, v := range slc {
	// 		if v <= min {
	// 			min = v
	// 		} else if v >= max {
	// 			max = v
	// 		}
	// 	}
	// 	fmt.Println(min, max)
	// }