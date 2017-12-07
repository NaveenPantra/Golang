package main

import "fmt"

func main() {
	//Interpreted String
	name := "Naveen Pantra"
	//Raw String
	fmt.Print(`
	<!DOCTYPE html>
	<html>
	<head>
	<meta charset="utf-8" language="en" />
	<title>Raw String</title>
	</head>
	<body>
	<h1>`, name, `</h1>
	</body>
	</html>
	`)
}

/*
-*- name := "Naveen Pantra"
    --> This is the interpreted string, i.e it cares about the all escape characters.
	--> This is made up of the rune literals so it cares about all the characters in the UTF-8.
    --> EX: change the back ticks in the print statement to the double quotes (shows the error because it is the Interpreted string)

-*- In the print statement the string is given in the back ticks it is called raw string
    --> Raw string do not care about the escaped characters in between the string.
    --> It only care about the back ticks(`) that should not appear in between the string.
 */