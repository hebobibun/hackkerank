package main

import (
	"fmt"
)

func angryProfessor(k int32, a []int32) string {
    // Write your code here
	var res string
	var count int

	for i := 0; i < len(a); i++ {
		if a[i] <= 0 {
			count++
		}
	}

	if count >= int(k) {
		res = "YES"
	} else {
		res = "NO"
	}

	return res
}

func main() {
	fmt.Println(angryProfessor(3, []int32{-2, -1, 0, 1, 2})) // YES
}