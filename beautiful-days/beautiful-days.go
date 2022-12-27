package main

import (
	"fmt"
	"strconv"
)

func reverseNum(num int) int32 {
	var res int

	numstr := strconv.Itoa(num)
	var result string

    for _, v := range numstr {
        result = string(v) + result
    }

	res, _ = strconv.Atoi(result)
	
	return int32(res)
}

func beautifulDays(i int32, j int32, k int32) int32 {
    // Write your code here
	var beautiful int32

	for z := i; z <= j; z++ {
		if (z - int32(reverseNum(int(z)))) % k == 0 {
			beautiful++
		}
	}

	return beautiful
}

func main() {
	fmt.Println(beautifulDays(20, 23, 6)) // 2
}