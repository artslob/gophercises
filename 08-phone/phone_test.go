package main

import (
	"fmt"
	"github.com/artslob/gophercises/08-phone/normalization"
	"testing"
)

func ExampleHello() {
	fmt.Println("hello")
	// Output: hello
}

var cases = []struct {
	input, expected normalization.Phone
}{
	{"", ""},
	{" ", ""},
	{
		"234567890",
		"234567890",
	},
	{
		"23 456 7891",
		"234567891",
	},
	{
		"(123) 456 7892",
		"1234567892",
	},
	{
		"(123) 456-7893",
		"1234567893",
	},
	{
		"23-456-7894",
		"234567894",
	},
	{
		"23-456-7890",
		"234567890",
	},
	{
		"234567892",
		"234567892",
	},
	{
		"(123)456-7892",
		"1234567892",
	},
	{
		" (123) 456 7892 ",
		"1234567892",
	},
}

func TestPhoneNormalize(t *testing.T) {
	for _, testCase := range cases {
		input := testCase.input
		expected := testCase.expected
		result := normalization.Normalize(input)
		if result != expected {
			t.Error("for", input, "normalization expected to be", expected, "got:", result)
		}
		resultLength := len(result)
		if result != "" && resultLength != 9 && resultLength != 10 {
			t.Error("result phone length should be 9 or 10, result:", result, "got:", resultLength)
		}
	}
}
