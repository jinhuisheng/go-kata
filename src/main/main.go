package main

import (
	"fmt"
)

func main() {
	if err := test(); err != nil {
		fmt.Println("Main Error:", err)
	}
}

func catchPanic(err *error) {
	if r := recover(); r != nil {
		fmt.Printf("%s : PANIC Defered : %v\n", r)
		if err != nil {
			*err = fmt.Errorf("%v", r)
		}
	}
}

func mimicError(key string) error {
	return fmt.Errorf("Mimic Error : %s", key)
}

func test() (err error) {
	defer catchPanic(&err)
	fmt.Println("Start Test")

	err = mimicError("1")

	panic("Mimic Panic")
}
