package main

import "fmt"

func main() {
	i := 0
	for {
		i++
		if i%2 == 0 {
			fmt.Printf("%d is even.\n", i)
			continue
		}
		fmt.Printf("%d is odd.\n", i)
		if i >= 50 {
			break
		}
	}
}

/*
 continue: This statement just skip the present(currently executing) iteration
 */
