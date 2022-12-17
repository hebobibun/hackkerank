package main

import (
	"fmt"
	"strings"
)

func countingValleys(steps int32, path string) int32 {
    // Write your code here
	var sum int32
	var count int32

	res := strings.Split(path, "")
	fmt.Println(res)

	for i := 0; i < len(res); i++ {
		if sum == -1 && res[i] == "U" {
			count += 1
		}

		if res[i] == "U" {
			sum += 1
		} else if res[i] == "D" {
			sum -= 1
		}
	}

	return count
}

func main() {
	fmt.Println(countingValleys(8, "UDDDUDUU")) // 1
}