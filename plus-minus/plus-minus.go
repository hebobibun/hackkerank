package main

import "fmt"

func plusMinus(arr []int32) {
    // Write your code here
	var neg float32
	var pos float32
	var zer float32

	for i := 0; i < len(arr); i++ {
		if arr[i] > 0 {
			pos++
		} else if arr[i] < 0 {
			neg++
		} else if arr[i] == 0 {
			zer++
		} 
	}

	pos = pos / float32(len(arr))
	neg = neg / float32(len(arr))
	zer = zer / float32(len(arr))

	fmt.Println(pos)
	fmt.Println(neg)
	fmt.Println(zer)
}

func main() {
	arr := []int32{-4, 3, -9, 0, 4, 1}
	plusMinus(arr) // 0.500000 0.333333 0.166667
}