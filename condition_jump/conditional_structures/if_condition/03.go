package main

func main() {
	if false {
		print("if")

	} else if true {
		print("else if")
	} else {
		print("else")
	}
}

/*
 -*- Golang follow strict typo
	 --> else if statement should not go to the newline if should start immediately after the if statement parentheses.
	 --> else statement should also start immediately after the closed parentheses.
 */
