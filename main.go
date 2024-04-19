package main

import "fmt"

func main() {
	fmt.Printf("%d \t %b\n", 1, 1)
	fmt.Printf("%d \t %b \n", 1<<1, 1<<1)
	fmt.Printf("%d \t %b \n", 2<<2, 2<<2)

	fmt.Printf("%d \t %b \n", 1>>1, 1>>1)
	fmt.Printf("%d \t %b \n", 2>>2, 2>>2)
}
