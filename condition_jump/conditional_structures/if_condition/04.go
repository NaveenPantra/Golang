package main

func main() {
	b := true
	if b {
		if b {
			print("Nested if")
		} else if !b {
			print("Nested else if")
		} else {
			print("Nested else")
		}
	} else {
		print("else statement")
	}
}

/*
 -*- This program shows the nested "if|else if|else" statements.
*/
