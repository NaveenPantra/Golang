package main

import "fmt"

var( b, c, d = 1, 2, 3)

const a string = "Naveen Pantra"

const (
	pi = 3.14
	e = 0.234
)

func main() {

	fmt.Printf("Value of constant A:  %c", '$')
	fmt.Printf("\nThe value of the variables B: %d, c: %d, d: %d", b, c, d)
	fmt.Printf("\nValue of constant pi: %v, e: %v", pi, e)
}
