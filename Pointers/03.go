package main

import "fmt"

func main() {
	var a int
	var b *int
	//var b *int = &a (One can do like this also)
	b = &a
	fmt.Printf("Enter the value: ")
	fmt.Scanf("%d", b)
	fmt.Printf("Value of a(a): %d, b(*b): %d", a, *b)
	fmt.Printf("\nEnter value to change: ")
	fmt.Scanf("%d", &a)
	fmt.Printf("The value of a(a): %d, b(*b): %d", a, *b)
	fmt.Print("\n\nThe address of a(b): ", b)
	fmt.Print("\n\nThe value at the address ", b, ", Is ", *b)
}

/*

 -*- The pointers are jusst like in c-language.
*/
