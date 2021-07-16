package fizzbuzz

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_FizzBuzz(t *testing.T) {
	cases := []struct {
		number   int
		expected string
	}{
		{1, "1"},
		{3, "Fizz"},
		{5, "Buzz"},
		{15, "FizzBuzz"},
	}
	for _, test := range cases {
		t.Run(fmt.Sprintf("input:%d,expected:%s", test.number, test.expected), func(t *testing.T) {
			assertResult(t, test.expected, test.number)
		})
	}
}

func Test_输入数字不能小于0(t *testing.T) {
	_, err := Of(-1)
	assert.Error(t, err, "输入数字不能小于0")
}

func assertResult(t *testing.T, expected string, number int) {
	result, err := Of(number)
	assert.Equal(t, expected, result)
	assert.Equal(t, nil, err)
}
