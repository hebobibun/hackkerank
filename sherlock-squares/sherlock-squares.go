package main

import (
	"fmt"
	"math"
)

func squares(a int32, b int32) int32 {
    // Write your code here
	A := int32(math.Sqrt(float64(a)))
	B := int32(math.Sqrt(float64(b)))

	res := B - A
	if A * A >= a {
		res++
	}

	return res
}

// func squares(a int32, b int32) int32 {
// 	// Write your code here
// 	var res int32

// 	for i := a; i <= b; i++ {
// 		if isPerfectSquare(int(i)) {
// 			res++
// 		}

// 		if i * i > b {
// 			break
// 		}
// 	}

// 	return res
// }

// func isPerfectSquare(no int) bool {
// 	int_root := int(math.Sqrt(float64(no)))
// 	return int_root * int_root == no
// }

func main() {
	fmt.Println(squares(24, 49)) // 3
	fmt.Println("======================")
	fmt.Println(squares(17, 24)) // 0
	fmt.Println("======================")
	fmt.Println(squares(1, 1000000)) // 1000
	
}


