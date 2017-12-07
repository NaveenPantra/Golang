package main

import "fmt"

func main() {
	i := 0
	for {
		if i > 10 {
			break //Unconditional break statement
		}
		fmt.Printf("i = %d\n", i)
		i++
	}
}

/*

 -*- break : This statement cause the program flow to just jump out the iteration.

*/
