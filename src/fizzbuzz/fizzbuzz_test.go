package fizzbuzz

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_输入数字小于0_抛出异常(t *testing.T) {
	_, err := Of(-1)
	assert.EqualErrorf(
		t,
		err,
		"输入数字不能小于0",
		err.Error(),
	)
}

func Test_FizzBuzz(t *testing.T) {
	cases := []struct {
		number   int
		Expected string
	}{
		{1, "1"},
		{3, "Fizz"},
		{5, "Buzz"},
		{15, "FizzBuzz"},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf(fmt.Sprintf("input: %d,expected:%s", tc.number, tc.Expected)), func(t *testing.T) {
			assertResult(t, tc.number, tc.Expected)
		})
	}
}

func assertResult(t *testing.T, number int, expected string) {
	result, _ := Of(number)
	assert.Equal(t, expected, result)
}
