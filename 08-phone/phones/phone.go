package phones

import "unicode"

type Phone string

func Normalize(phone Phone) Phone {
	var result []rune
	for _, r := range phone {
		if unicode.IsDigit(r) {
			result = append(result, r)
		}
	}
	return Phone(result)
}
