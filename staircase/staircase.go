package main

import "fmt"

func staircase(n int32) {
    // Write your code here
	for i := int32(0); i < n; i++ {
		for j := int32(0); j < n; j++ {
			if i + j < n-1 {
				fmt.Print(" ")
			} else if i + j >= n -1 {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}

func main() {
	staircase(6)
	// 	    #
	//     ##
	//    ###
	//   ####
	//  #####
	// ######
}