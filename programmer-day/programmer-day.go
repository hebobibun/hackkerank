package main

import "fmt"

func dayOfProgrammer(year int32) string {
    // Write your code here
	var theDay string

	if year <= 1917 {
		if year % 4 == 0 {
			theDay = fmt.Sprintf("%d.09.%v", 256-(31+29+31+30+31+30+31+31), year)
		} else {
			theDay = fmt.Sprintf("%d.09.%v", 256-(31+28+31+30+31+30+31+31), year)
		}

	} else if year == 1918 {
		theDay = fmt.Sprintf("%d.09.%v", 256-(31+15+31+30+31+30+31+31), year)
	} else {
		if (year % 4 == 0 && year % 100 != 0) || year % 400 == 0 {
			theDay = fmt.Sprintf("%d.09.%v", 256-(31+29+31+30+31+30+31+31), year)
		} else {
			theDay = fmt.Sprintf("%d.09.%v", 256-(31+28+31+30+31+30+31+31), year)
		}
	}

	return theDay
}

func main() {
	fmt.Println(dayOfProgrammer(1917)) // 13.09.1917
	fmt.Println(dayOfProgrammer(2017)) // 13.09.1987
}