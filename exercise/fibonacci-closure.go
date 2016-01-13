package main

import "fmt"

func fibonacci() func() int {
	var i = 0
	var j = 1
	return func() int {
		j += i
		i = j - i
		return i
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
