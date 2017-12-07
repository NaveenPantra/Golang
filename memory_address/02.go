package main

import "fmt"

// I think this is wrong.
const foot float32 = 3.2

func main() {
	var a float32
	fmt.Printf("Enter Meters: ")
	fmt.Scanf("%f", &a)
	//fmt.Scan(&a);
	var b = a * foot
	fmt.Printf("%.2f meters is %.2f foots", a, b)
}

/*

-*- In Formatted printf "%.2f",this .2 indicate that only 2 decimal point should go to the stdout.
    --> More on this in the c-language Classwork.
*/
