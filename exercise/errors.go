package main

import (
	"fmt"
)

type MyError struct {
}

func (e *MyError) Error() string {
	return "Parameter should greater than 0"
}

func Sqrt(x float64) (float64, error) {
	if x < 0.0 {
		return 0, &MyError{}
	}
	z := 1.0
	for i := 0; i < 10; i++ {
		z = (z + x/z) / 2
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
