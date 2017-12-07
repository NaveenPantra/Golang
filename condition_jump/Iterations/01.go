package Iterations

import "fmt"

func main() {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Printf("%d\t", j)
		}
		//fmt.Printf("i: %d", i)
		fmt.Println()
	}
}
