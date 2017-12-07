package _switch

import "fmt"

func main() {
	var a int
	a = 12
	switch a.(type) { 	//This is called assertion in the golang
	case int8, int16, int32, int64, uint:
		fmt.Printf("It's INTEGER.")
	case rune:
		fmt.Printf("It's RUNE.")
	case string:
		fmt.Printf("It's STRING.")
	case float32, float64 :
		fmt.Printf("It's FLOAT")
	default:
		fmt.Printf("Cannot fetch bro..!")
	}
}
/*
 -*- In golang to know the type of the variable or the identifier we can find by using
	 --> <identifier>.(type) this can only be used in the switch statement not outside of that.
	 --> Is this is used outside the switch statement the compiler flag an error.

 -*- Programmer can do more cool stuff on the assertion by using the data-structures like the "struct" wait to see.it..
 */
