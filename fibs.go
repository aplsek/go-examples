package main

import (
	"fmt"
	
)

//F(0) = 0; F(1) = 1; F(n) = F(n–1) + F(n–2) 

var fibs map[int]int


func fib(i int) int {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	if v, ok := fibs[i]; ok {
		return v
	}
	i1 := fib(i-1)
	i2 := fib(i-2)
	fibs[i] = i1 + i2
	return i1+i2
}

func main() {
	//n := "12"
	fibs = make(map[int]int,0)
	//f := fib(10)
	
	fmt.Printf("Hello, playground: %v\n", fib(10))
	fmt.Printf("Hello, playground: %v\n", fib(11))
	fmt.Printf("Hello, playground: %v\n", fib(9))
}
