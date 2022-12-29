package main

import "fmt"

func timeInWords(h int32, m int32) string {
    // Write your code here
	var res string

	numbers := make(map[int32]string)
	numbers[1] = "one"
	numbers[2] = "two"
	numbers[3] = "three"
	numbers[4] = "four"
	numbers[5] = "five"
	numbers[6] = "six"
	numbers[7] = "seven"
	numbers[8] = "eight"
	numbers[9] = "nine"
	numbers[10] = "ten"
	numbers[11] = "eleven"
	numbers[12] = "twelve"
	numbers[13] = "thirteen"
	numbers[14] = "fourteen"
	numbers[15] = "quarter"
	numbers[16] = "sixteen"
	numbers[17] = "seventeen"
	numbers[18] = "eighteen"
	numbers[19] = "nineteen"
	numbers[20] = "twenty"
	numbers[21] = "twenty one"
	numbers[22] = "twenty two"
	numbers[23] = "twenty three"
	numbers[24] = "twenty four"
	numbers[25] = "twenty five"
	numbers[26] = "twenty six"
	numbers[27] = "twenty seven"
	numbers[28] = "twenty eight"
	numbers[29] = "twenty nine"
	numbers[30] = "half"

	res = numbers[h]

	if m == 0 {
		res = numbers[h] + " " + "o' clock"
		fmt.Println("== 0")
	}
	
	if m > 0 && m <= 30 {
		if m == 1 {
			res = numbers[m] + " " + "minute" + " " + "past" + " " + numbers[h]
		} else if m == 15 || m == 30 {
			res = numbers[m] + " " + "past" + " " + numbers[h]
		} else {
			res = numbers[m] + " " + "minutes" + " " + "past" + " " + numbers[h]
		}	

		fmt.Println("< 30")
	}

	if m > 30 {
		if m == 45 {
			res = numbers[60-m] + " " + "to" + " " + numbers[h+1]
		} else {
			res = numbers[60-m] + " " + "minutes" + " " + "to" + " " + numbers[h+1]
		}
		fmt.Println("> 30")
	}

	return res
}

func main() {
	fmt.Println(timeInWords(5, 00)) // five o' clock
	fmt.Println(timeInWords(5, 01)) // one minute past five
	fmt.Println(timeInWords(5, 10)) // ten minutes past five
	fmt.Println(timeInWords(5, 15)) // quarter past five
	fmt.Println(timeInWords(5, 30)) // half past five
	fmt.Println(timeInWords(5, 40)) // twenty minutes to six
	fmt.Println(timeInWords(5, 45)) // quarter to six
	fmt.Println(timeInWords(5, 47)) // thirteen minutes to six
	fmt.Println(timeInWords(5, 28)) // twenty eight minutes past five
}