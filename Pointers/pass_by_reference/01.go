package main

import "fmt"

func zero(z *int) {
	fmt.Printf("\nThe address of z: %p", z)
	fmt.Printf("\nThe value of z: %d", *z)
	*z = 0
	fmt.Printf("\nThe z is assigned to 0")
}

func main() {
	var z int
	fmt.Printf("Enter the value: ")
	fmt.Scanf("%d", z)
	fmt.Printf("Passing the value to the zero function(pass_by_reference)")
	zero(&z)


}

/*
 -*- %p in the printf print the base 16 notation leading with 0x.
 */