package main

import (
	"fmt"
	"math"
)

func catAndMouse(x int32, y int32, z int32) string {
	var res string

	xDistance := math.Abs(float64(x)-float64(z))
	yDistance := math.Abs(float64(y)-float64(z))

	if xDistance == yDistance {
		res = "Mouse C"
	} else if xDistance > yDistance {
		res = "Cat B"
	} else {
		res = "Cat A"
	}

	return res
}

func main() {
	fmt.Println(catAndMouse(1, 2, 3)) // CAT B
	fmt.Println(catAndMouse(1, 3, 2)) // Mouse C
}