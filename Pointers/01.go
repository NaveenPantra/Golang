package main

import "fmt"

func main() {
	var j *int
	var k **int
	k = &j
	fmt.Printf("Address of j: %p, Address of k: %p", &j, &k)

}
