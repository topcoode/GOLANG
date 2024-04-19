package main

import (
	"fmt"
	"time"
)

func main() {
	/*
	   \r: Represents a carriage return character. Moves the cursor to the beginning of the current line.
	*/
	/*
		fmt.Print("Progress: ")
		fmt.Print("10%")
		time.Sleep(time.Second * 2)
		fmt.Print("\r") // Move cursor to the beginning of the line
		fmt.Print("Progress: ")
		fmt.Print("20%")
	*/
	/*
		\b: Represents a backspace character. It moves the cursor one position backward.
	*/
	/*
		fmt.Print("Loading |")
		time.Sleep(time.Second)
		fmt.Print("\b/") // Move cursor back and overwrite character
		time.Sleep(time.Second)
		fmt.Print("\b-") // Move cursor back and overwrite character
		time.Sleep(time.Second)
		fmt.Print("\b\\") // Move cursor back and overwrite character
		time.Sleep(time.Second)
		fmt.Print("\b|") // Move cursor back and overwrite character
		time.Sleep(time.Second)
		fmt.Println("\nDone!")
	*/

	/*
		\f: Represents a form feed character. It moves the cursor to the beginning of the next page or form.

	*/
	fmt.Println("Page 1")
	time.Sleep(time.Second * 5)
	fmt.Println("\f")
	fmt.Println("Page 2")

}
