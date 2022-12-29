package main

import (
	"fmt"
	"strings"
)

func marsExploration(s string) int32 {
    // Write your code here
	var res int32
	splitS := strings.Split(s, "")
	fmt.Println(splitS)

	slc := []string{}

	for i := 0; i < len(splitS); i++ {
		if len(slc) == 0 {
			if splitS[i] != "S" {
				res++
			}
			slc = append(slc, splitS[i])
		} else if len(slc) == 1 {
			if splitS[i] != "O" {
				res++
			}
			slc = append(slc, splitS[i])
		} else if len(slc) == 2 {
			if splitS[i] != "S" {
				res++
			}
			slc = append(slc, splitS[i])
		}
		
		if len(slc) == 3 {
			slc = []string{}
		}
	}

	return res
}

func main() {
	fmt.Println(marsExploration("SOSSPSSQSSOR")) // 3
	fmt.Println(marsExploration("SOSSOSSOS")) // 0
	fmt.Println(marsExploration("SOSSOT")) // 1
}