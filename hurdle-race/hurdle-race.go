package main

import "fmt"

func hurdleRace(k int32, height []int32) int32 {
    // Write your code here
	var dos int32
	lim := k

	for i := 0; i < len(height); i++ {
		if height[i] > lim {
			dos = height[i] - k
			lim = height[i]
		}
	}

	return dos
}

func main() {
	fmt.Println(hurdleRace(4, []int32{1, 6, 3, 5, 2})) // 2
	fmt.Println(hurdleRace(4, []int32{1, 5, 3, 5, 6})) // 2
	fmt.Println(hurdleRace(3, []int32{1, 3, 3, 1, 4})) // 1
	fmt.Println(hurdleRace(1, []int32{10, 6, 3, 9, 2})) // 9
	fmt.Println(hurdleRace(7, []int32{2, 5, 4, 5, 2})) // 0
	fmt.Println(hurdleRace(7, []int32{6, 5, 6, 5, 2})) // 0
	fmt.Println(hurdleRace(0, []int32{6, 5, 6, 5, 2})) // 6
}