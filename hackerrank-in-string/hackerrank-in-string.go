package main

import (
	"fmt"
	"strings"
)

func hackerrankInString(s string) string {
    // Write your code here
	res := "NO"

	hrank := map[int]string {
    0: "h",
	1: "a",	
	2: "c",	
	3: "k",	
	4: "e",	
	5: "r",	
	6: "r",	
	7: "a",	
	8: "n",	
	9: "k",
}

	splitS := strings.Split(s, "")
	slc := []string{}

	for i := 0; i < len(splitS); i++ {
		if splitS[i] == hrank[len(slc)] {
			slc = append(slc, splitS[i])
		}

		if len(slc) == 10 {
			break
		}
	}

	joinS := strings.Join(slc, "")

	if joinS == "hackerrank" {
		res = "YES"
	}

	return res
}

func main() {
	fmt.Println(hackerrankInString("hereiamstackerrank")) // YES
	fmt.Println(hackerrankInString("hackerworld")) // NO
}