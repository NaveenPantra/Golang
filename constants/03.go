package main

import "fmt"

const (
	iota_1 = iota
	iota_2 = iota
	iota_3 = iota
)

const (
	iota_4 = iota
	iota_5
	iota_6
)

const iota_7 = iota
const iota_8 = iota

func main() {
	fmt.Printf("The vales of \"IOTA\"")
	fmt.Printf("\niota_1: %v\niota_2: %v\niota_3: %v\n\niota_4: %v\niota_5: %v\niota_6: %v\n\niota_7: %v\niota_8: %v\n", iota_1, iota_2, iota_3, iota_4, iota_5, iota_6, iota_7, iota_8)
}

/*
 ioat: it is the third letter in the greek alphabet, It means the smallest number possible.

 -*- In the first const declaration the iota get increased from 0 to 2
 -*- In the second const declaration the single iota is given and it will assume that all the constants of type iota so it will increase from 0 to 2
 -*- In the third (this is just a simple declaration)
 -*- In t
 */