package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostfixToPrefix(t *testing.T) {
	examples := map[string]string{
		"2 45 + 6 + 24 +":            "+ + + 2 45 6 24",
		"60000321 764 +":             "+ 60000321 764",
		"26 89 13 +":                 "too many operands",
		"98 4 * 234 12 * +":          "+ * 98 4 * 234 12",
		"hello expression":           "too many operators",
		"9.007 765.9999994 + 56 ^ /": "too many operators",
		"":                           "invalid input expression",
		"4 2 - 3 * 5+":               "+ * - 4 2 3 5",
	}

	for postfix, expected := range examples {
		res, err := PostfixToPrefix(postfix)
		if assert.Nil(t, err) {
			assert.Equal(t, expected, res)
		}
	}
}

func ExamplePostfixToPrefix() {
	res, err := PostfixToPrefix("34 5 +")
	if err != nil {
		panic(err)
	} else {
		fmt.Println(res)
	}

	// Output:
	// + 34 5
}
