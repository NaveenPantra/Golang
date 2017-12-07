package main

import "fmt"

func zero(x int) {
	//x := 0 //Shows error (Try it)
	fmt.Printf("\nNow the x(Variable) is assigned to 0 in \"zero function\"")
	x = 0
}

func main() {
	var x int
	fmt.Printf("Enter the value: ")
	fmt.Scanf("%d", &x)
	fmt.Printf("Passing the x(variable) to zero function")
	zero(x)
	fmt.Printf("\nThe value of x is still %d", x)
}

/*
 -*- This program demonstrate that, In the go-programming language every thing will work on the "pass by value".
 -*- If the variable that was already declared got assigned with " := " operator then it is an error.
     --> " := " this operator is only for the new variables (Initialization).
*/
