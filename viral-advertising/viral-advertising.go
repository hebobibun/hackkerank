package main

import "fmt"

func viralAdvertising(n int32) int32 {
    // Write your code here
	var numLike int
	var like int

	for i := 1; i <= int(n); i++ {
		if i == 1 {
			numLike = 5 / 2
		} else if i == 2 {
			like = (numLike * 3) / 2
			numLike += like
		} else if i > 2 {
			like = (like * 3) / 2
			numLike += like
		}
		
	}

	return int32(numLike)
}

func main() {
	fmt.Println(viralAdvertising(3)) // 9
}