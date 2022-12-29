package main

import (
	"fmt"
	"unicode"
)

func pangrams(s string) string {
    // Write your code here
	res := "not pangram"

	alphabet := map[rune]bool{}
    for _, r := range "abcdefghijklmnopqrstuvwxyz" {
        alphabet[r] = true
    }

	seen := map[rune]bool{}
    for _, r := range s {
        r = unicode.ToLower(r)
        if alphabet[r] {
            seen[r] = true
        }
    }

    allSeen := true
    for r := range alphabet {
        if !seen[r] {
            allSeen = false
            break
        }
    }

    if allSeen {
        res = "pangram"
    }

	return res
}

func main() {
	fmt.Println(pangrams("We promptly judged antique ivory buckles for the next prize")) // pangram
	fmt.Println(pangrams("We promptly judged antique ivory buckles for the prize")) // not pangram
}