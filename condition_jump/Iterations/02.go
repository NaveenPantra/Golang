package Iterations

import "fmt"

func main() {
	i := 0
	fmt.Println("2 Table: ")
	for i <= 10 {
		fmt.Printf("2 * %d = %d", i, (2 * i))
		i++
		fmt.Println()
	}
}
