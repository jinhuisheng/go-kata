package fizzbuzz

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_普通数字_返回本身(t *testing.T) {
	assertResult(t, 1, "1")
}

func Test_被3整除_返回Fizz(t *testing.T) {
	assertResult(t, 3, "Fizz")
}

func Test_被5整除_返回Buzz(t *testing.T) {
	assertResult(t, 5, "Buzz")
}

func Test_同时被3和5整除_返回FizzBuzz(t *testing.T) {
	assertResult(t, 15, "FizzBuzz")
}

func assertResult(t *testing.T, number int, expected string) {
	result, _ := Of(number)
	assert.Equal(t, expected, result)
}

func Test_输入数字小于0_抛出异常(t *testing.T) {
	_, err := Of(-1)
	assert.EqualErrorf(
		t,
		err,
		"输入数字不能小于0",
		err.Error(),
	)
}

func TestAdd(t *testing.T) {
	cases := []struct{ A, B, Expected int }{
		{1, 1, 2},
		{1, -1, 0},
		{1, 0, 1},
		{0, 0, 0},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("%d + %d", tc.A, tc.B), func(t *testing.T) {
			actual := tc.A + tc.B
			if actual != tc.Expected {
				t.Fatal("failure")
			}
		})
	}
}
