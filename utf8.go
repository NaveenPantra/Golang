package main

import "fmt"

func main() {
	i := 0
	for {
		fmt.Printf("%c\n", i)
		i++
		if i == 500 {
			break
		}
	}
}
