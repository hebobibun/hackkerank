package main

import "fmt"

func getMoneySpent(keyboards []int32, drives []int32, b int32) int32 {
	var res int32

    for i := 0; i < len(keyboards); i++ {
        for j := 0; j < len(drives); j++ {
            if keyboards[i] + drives[j] > b && res == 0 {
                res = -1
            } else if keyboards[i] + drives[j] > res && keyboards[i] + drives[j] <= b {  
                res = keyboards[i] + drives[j]
            } 
        }
    }

    return res
}


func main() {
	fmt.Println(getMoneySpent([]int32{40, 50, 60}, []int32{5, 8, 12}, 60)) // 58
	fmt.Println(getMoneySpent([]int32{3, 1}, []int32{5, 2, 8}, 10)) // 9
	fmt.Println(getMoneySpent([]int32{4}, []int32{5}, 5)) // -1
	fmt.Println(getMoneySpent([]int32{4}, []int32{1, 5}, 5)) // 5
}