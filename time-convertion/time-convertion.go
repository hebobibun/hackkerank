package main

import (
	"fmt"
	"strconv"
	"strings"
)

func timeConversion(s string) string {
    // Write your code here
	slc := strings.Split(s, ":")
	hourToint, _ := strconv.Atoi(slc[0])

	if strings.Contains(s, "AM") {
		if hourToint == 12 {
			hourToint -= 12
		}
	} else if strings.Contains(s, "PM") {
		if hourToint != 12 {
			hourToint += 12
		} 
	}

	hourToStr := strconv.Itoa(hourToint)

	if len(hourToStr) == 0 {
		hourToStr += "0"
	} else if len(hourToStr) < 2 {
		hourToStr = "0" + hourToStr
	}

	slc[0] = hourToStr

	res := strings.Join(slc, ":")
	res = strings.ReplaceAll(res, "PM", "")
	res = strings.ReplaceAll(res, "AM", "")

	return res
}

func main() {
	fmt.Println(timeConversion("07:05:45PM")) // 19:05:45
	fmt.Println(timeConversion("11:05:45PM")) // 19:05:45
	fmt.Println(timeConversion("01:05:45PM")) // 19:05:45
	fmt.Println(timeConversion("12:05:45PM")) // 19:05:45
	fmt.Println(timeConversion("12:05:45AM")) // 00:05:45
	fmt.Println(timeConversion("01:05:45AM")) // 01:05:45
	fmt.Println(timeConversion("02:05:45AM")) // 02:05:45
	fmt.Println(timeConversion("09:05:45AM")) // 02:05:45
	fmt.Println(timeConversion("10:05:45AM")) // 02:05:45
	fmt.Println(timeConversion("11:05:45AM")) // 02:05:45

}