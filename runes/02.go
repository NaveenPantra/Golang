package main

import "fmt"

func main() {
	for i := 3072; i < 3199; i++ {
		//j = byte(i)
		fmt.Printf("UTF-8\t%d\tis\t%c\t", i, i)
		fmt.Printf("With bytes %x\n", []byte(string(i)))
		//fmt.Println("With bytes ", []byte(string(i)))
	}

}

/*
-*- This code print all the telugu font characters in UTF-8 code format.
-*- In the println statement
    --> This give the output in terms of the bytes it using.(Telugu font uses 3 bytes to represent the single character).
    --> []byte This is conversion of the String that(String) actually converting int32 'i' to String and finally to the byte(by []byte).
*/
