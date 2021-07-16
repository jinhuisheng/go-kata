package fizzbuzz

import (
	"errors"
	"strconv"
)

func Of(number int) (result string, err error) {
	if number < 0 {
		err = errors.New("输入数字不能小于0")
	}
	if number%3 == 0 {
		result += "Fizz"
	}
	if number%5 == 0 {
		result += "Buzz"
	}
	if result == "" {
		result = strconv.Itoa(number)
	}
	return
}
