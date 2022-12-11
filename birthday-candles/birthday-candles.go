package main

func birthdayCakeCandles(candles []int32) int32 {
    // Write your code here
	// var tallest []int32
	var tmp int32
	for _, v := range candles {
		if v >= tmp {
			tmp = v
		}
	}

	var tmp2 int32

	for _, v := range candles {
		if v == tmp {
			tmp2++
		}
	}

	return tmp2
}

func main() {
	birthdayCakeCandles([]int32{1,4,2,5,1,5,3}) // 2
}	