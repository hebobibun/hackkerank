package main

import "fmt"

func gradingStudents(grades []int32) []int32 {
    // Write your code here
	for i := int32(40); i <= 100; i+=5 {
		for j := 0; j < len(grades); j++ {
			if grades[j] >= 38 && i > grades[j] {
				if i - grades[j] < 3 {
					grades[j] = i
				}
			}
		}
	}
		
	return grades
}

func main() {
	fmt.Println(gradingStudents([]int32{73, 67, 38, 33})) // 75, 67, 40, 33
	fmt.Println(gradingStudents([]int32{88, 62, 36, 98})) // 75, 70, 40, 23
}
