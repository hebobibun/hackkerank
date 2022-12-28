package main

import "fmt"

func libraryFine(d1 int32, m1 int32, y1 int32, d2 int32, m2 int32, y2 int32) int32 {
    // Write your code here
	var res int32

	if d1 <= d2 && m1 <= m2 && y1 <= y2 {
		res = 0
	} else if d1 > d2 && m1 == m2 && y1 == y2 {
		res = (d1 - d2) * 15
	} else if m1 > m2 && y1 == y2 {
		res = (m1 - m2) * 500
	} else if y1 > y2 {
		res = 10000
	}

	return res
}

func main() {
	fmt.Println(libraryFine(9, 6, 2015, 6, 6, 2015)) // 45
	fmt.Println(libraryFine(1, 1, 2018, 30, 12, 2017)) // 10000
}