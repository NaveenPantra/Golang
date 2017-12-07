package main

import "fmt"

func assign() func() int {
	var x = 0
	return func() int {
		x++;
		return x
	}
}

func main() {
	increment := assign()
	fmt.Print("Value of the X: ", increment())
	fmt.Print("\nValue of the X: ", increment())
}

/*
Closure: The function that has the accessability of any variable can hold the variable even the function was returned by the top function,
	this because the compiler maintain the scope of every function and and hold them even the top function was returned.

 */