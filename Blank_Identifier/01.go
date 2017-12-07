package main

import "fmt"

func main() {
	//c := 1  --> This the blank identifier.
	b := 2
	fmt.Printf("The value of B:%v", b);
}

/*
	The use of the blank identifier is that it does not allow the program to do not declare the unused variables.
	The variable 'c' is declared but not used in the program so, code will give an error.
 */