package main

import (
	"fmt"
)

//This is the package level scope
var Fullname string = "Naveen Pantra"

//This are all package scope
//This are also package level scope
//This called Initilization (i.e Declaring and assigning the variables at the same time)
var fname string = "Naveen"
var sec string = "b"

//This variables are blank Identifiers (i.e nothing is inserted in the variables), also called declararion.
var lname string
var dob string
var branch string

func main() {
	//This age, college variable is directly initilized with 19 and it is of type INTEGER.
	//The age, college variable is function scope and cannot be used out side the function or block.
	age := 19
	college := `Sree Vidyanikethan Engineering College`

	//This are the assignmen for the variables that are daclared on the package scope
	lname = `Pantra`
	dob = "20-05-1998"
	branch = "Information Technology"
	fmt.Printf("\n%s %s  studying %s-%s in %s.\n", fname, lname, branch, sec, college)
	fmt.Printf("\nHis data of birth: %s and his age up to now is: %d\n", dob, age)
}
