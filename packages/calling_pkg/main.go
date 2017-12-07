package main

import (
	"fmt"
	"github.com/NaveenPantra/Golang/packages/called_pkg"
	//The called package contain the three go files
	// 1.name.go (This contain the Name).
	// 2.reverse.go (contatin Reverse() visible to outside package, The Reverse() call the recerse() in the same package).
	// 3. reverse_two.go (contain reverse() not visible to outside the package, This was called by reverse.go).
)

func main() {
	fmt.Print("Hello World \n")
	// call any name or method form other package by using the package name(folder name).
	// Here called_pkg is  the package name(folder name).
	fmt.Print("Name: " + called_pkg.Name)
	fmt.Print("\nReverse Name: " + called_pkg.Reverse())
}
