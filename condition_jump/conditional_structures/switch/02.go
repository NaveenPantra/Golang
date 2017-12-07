package _switch

import "fmt"

func main() {
	name := "Naveen Pantra"
	switch  {
	case len(name) == 13 :
		fmt.Printf("Your are my friend with name length: %d", len(name))
		fallthrough
	case name == "Naveen Pantra", name == "Naveen", name == "Pantra":
		fmt.Printf("\nYour Name is %s", name)
	default:
		fmt.Printf("\nNot catching you dude..")
	}
}


/*
 -*- In golang programmer can use the empty expression on the switch statement.
 -*- Then all the stuff inside the switch act as the same as the normal swich.
 */