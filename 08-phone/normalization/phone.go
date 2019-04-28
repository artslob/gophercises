package normalization

import "unicode"

type Phone string

func (phone Phone) Normalize() Phone {
	var result []rune
	for _, r := range phone {
		if unicode.IsDigit(r) {
			result = append(result, r)
		}
	}
	return Phone(result)
}
