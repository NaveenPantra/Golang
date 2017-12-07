package main

import "fmt"

const (
	_  = iota             //iota = 0 (This is a blank identifier just throw away the iota = 0)
	KB = 1 << (iota * 10) //iota = 1
	MB = 1 << (iota * 10) //iota = 2
	GB = 1 << (iota * 10) //iota = 3
	TB = 1 << (iota * 10) //iota = 4
)

func main() {
	fmt.Printf("The cosntants: ")
	fmt.Printf("\nKB: %b \t %d", KB, KB)
	fmt.Printf("\nMB: %b \t %d", MB, MB)
	fmt.Printf("\nGB: %b \t %d", GB, GB)
	fmt.Printf("\nTB: %b \t %d", TB, TB)

}

/*
 -*- Here the blank identifier just throw away the iota = 0 value by not storing it
     --> This is just like * that is in the formatted "scanf" in c-Language (Just read the value the value but do not store it)
         --> ex: scanf("%d%d%d", &a,&*b,&c) //This should be verified.
 			 --> Here the second value just read but do not store any where.

 -*- <<: is the bitwise left operator this shift the bits to the left position.
		 --> meaning: shift 1 left with number of zeros specified on the right side.
*/
