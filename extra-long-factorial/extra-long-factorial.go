package main

import (
	"fmt"
	"math/big"
)

func extraLongFactorials(n int32) {
    // Write your code here
	res := big.NewInt(1)
    for i := 2; i <= int(n); i++ {
        res.Mul(res, big.NewInt(int64(i)))
    }
    fmt.Println(res.String())
}

func main() {
	extraLongFactorials(25)
}