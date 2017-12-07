package main

import "fmt"

func main() {
	a := 'a'
	//b : ='z'
	for a <= 'z' {
		fmt.Printf("%c ", a)
		//if a == 'z' {
		//	break
		//}
		a++
	}
	a--
	fmt.Println()
	for a >= 'a' {
		fmt.Printf("%c ", a)
		a--
	}
}
