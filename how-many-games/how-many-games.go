package main

import "fmt"

func howManyGames(p int32, d int32, m int32, s int32) int32 {
    // Return the number of games you can buy
	var res int32
	var dec int32
	var count int32

	for res < s {
		if p > s {
			break
		} else if res == 0 && p <= s {
			res = p
			dec = p
			count++
			fmt.Println("if 1 : ", count, "dec : ", dec)
		} else if res <= s {
			if dec - d <= m {
				if res + m > s {
					break
				} else {
					res = res + m
					count++
					fmt.Println("if 2 : ", count, "dec : ", dec)
				}
				
			} else {
				if res + (dec - d) > s {
					break
				} else {
					res = res + (dec - d)
					dec = dec - d
					count++
					fmt.Println("if 3 : ", count, "dec : ", dec)
				}
				
			}
		}
	}

	return count
}

func main() {
	fmt.Println(howManyGames(20, 3, 6, 80)) // 6
	fmt.Println(howManyGames(20, 3, 6, 85)) // 7
	fmt.Println(howManyGames(100, 1, 1, 99)) // 0
	fmt.Println(howManyGames(100, 19, 1, 180)) // 1
}