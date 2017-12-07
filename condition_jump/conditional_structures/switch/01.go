package _switch

import "fmt"

func main() {
	i := 1
	switch i {
	case 1 | 0:		//Now this apply bitwise and on the 1, 0 an result to 1.
		fmt.Print("Hello dude..")
	default:
		fmt.Print("Nothing, gotta go bye!")
	}

	c := 'N'

	switch c {		//this also apply the bitwise-or operation he c, d and result to 103.
	case 'c' | 'd':
		fmt.Print("\nName start with C or D")
	case 'N' , 'n':
		fmt.Print("\nName start with N or n ")
	default:
		fmt.Print("\nCan't guess the name dude.")

	}

	j := "Naveen"

	switch j {
	case "Naveen", "Pantra":
		fmt.Printf("\nHai Naveen Pantra")
		fallthrough
	case "Hello":
		fmt.Print("\nHello.")

	default:
		fmt.Printf("\nSee u dude next time..")
	}
}

/*

 -*- In golang the switch statement just look like the switch statement in c-language.
 -*- Break is no needed in the switch in the golang, compiler put the default break statement.
 -*- you can explicitly can place a break statement, well that has no meaning.
 -*- instead of "break" statement, programmer can place an explicit statement called "fallthrough"
	 --> "fallthrough" statement will override the "break" statement
	 --> "fallthrough" statement is just like the with out the "break" statement in other languages.
	 --> "fallthrough" statement do not check the case condition just execute the next case program statement's.
 -*- Want to place more options in the case statement separate the conditions by using the comma(,).
	 --> If programmer try the bitwise-or(|) it just do the operation on the Integrals(Cannot apply on strings).
 -*- Programmer can use the logical-and operator on the case to build the conditions.
*/
