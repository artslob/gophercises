package main

import (
	"fmt"
	"unicode"
)

func camelcase(s string) int {
	count := 1
	for _, r := range s {
		if unicode.IsUpper(r) {
			count++
		}
	}
	return count
}

func caesarCipher(s string, k int32) string {
	var buffer []rune
	for _, r := range s {
		if unicode.IsLetter(r) {
			var base int32
			if unicode.IsLower(r) {
				base = 'a'
			} else {
				base = 'A'
			}
			r = base + (r-base+k)%26
		}
		buffer = append(buffer, r)
	}
	return string(buffer)
}

func main() {
	fmt.Println(camelcase("saveChangesInTheEditor"))
	fmt.Println(caesarCipher("middle-Outz", 2))
}
