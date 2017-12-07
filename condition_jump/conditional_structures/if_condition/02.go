package main

func main() {
	if name := "Naveen Pantra"; len(name) == 13 {
		print(name, " length is: ", len(name))
	}
	//print(name, " length is: ",len(name)) --> Scope of the variable "name".
}

/*
 -*- Programmer can do initialization before checking the condition in the if statement, this in go called initialization statement.
 -*- This hack is used in opening files and doing some stuff on that.
	 --> Explore "effective go" document in this initialization statement.
	 --> EX: in c-lang we use FILE pointer and check for the file existence(i.e: fp == NULL).
 -*- The scope of the variable "name" is in the if statement block only.
 -*-
*/
