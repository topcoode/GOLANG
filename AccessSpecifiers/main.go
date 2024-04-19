package main

import "fmt"

type Structure struct {
	Name  string
	TaxNo int
}

func main() {
	/*
		%v	the value in a default format

	*/
	var data int = 54
	fmt.Printf("the value of the v  %v", data)

	//when printing structs, the plus flag (%+v) adds field names
	structure := Structure{
		Name:  "Soarg",
		TaxNo: 545468723658,
	}
	fmt.Printf("using the \n%+v", structure)

	//%#v	a Go-syntax representation of the value

	fmt.Println("\n-------------------------------------------------------------->")
	var integer int = 56
	var integerdata int32 = 56
	var stringdata string = "string data"
	var booleandata bool = true

	fmt.Printf("\n structure : %#v", structure)
	fmt.Printf("\n integer : %#v", integer)
	fmt.Printf("\n integer : %#v", integerdata)
	fmt.Printf("\n string : %#v", stringdata)
	fmt.Printf("\n boolea : %#v", booleandata)
	fmt.Println("\n-------------------------------------------------------------->")
	// %T	a Go-syntax representation of the type of the value
	var a int
	var b float64
	var c string

	fmt.Printf("Type of a: %T\n", a) // Output: Type of a: int
	fmt.Printf("Type of b: %T\n", b) // Output: Type of b: float64
	fmt.Printf("Type of c: %T\n", c) // Output: Type of c: string
	fmt.Println("\n-------------------------------------------------------------->")
	//%%	a literal percent sign; consumes no value

	fmt.Printf("10%% discount\n") // Output: 10% discount

	fmt.Println("\n-------------------------------------------------------------->")
	//%t	the word true or false

	var a1 bool = true
	var b1 bool = false

	fmt.Printf("a is %t\n", a1) // Output: a is true
	fmt.Printf("b is %t\n", b1) // Output: b is false

	fmt.Println("\n-------------------------------------------------------------->")

	//%b	base 2
	/*
			binary numbers are expressed using a base-2 numeral system. In a base-2 system,
			each digit in a number can have one of two possible values: 0 or 1.

		   the %b format specifier is used with the
		   Printf family of functions to format and print integers in binary representation.
	*/
	var a2 int = 10
	var b2 int = 5

	fmt.Printf("a in binary: %b\n", a2) // Output: a in binary: 1010
	fmt.Printf("b in binary: %b\n", b2) // Output: b in binary: 101

	fmt.Println("\n-------------------------------------------------------------->")

	//%c	the character represented by the corresponding Unicode code point

	/*
		the %c format specifier is used with the Printf family of functions to format and print characters.


	*/

	var ch1 rune = 'A'
	var ch2 byte = 'B'

	fmt.Printf("Character 1: %c\n", ch1) // Output: Character 1: A
	fmt.Printf("Character 2: %c\n", ch2) // Output: Character 2: B

	fmt.Println("\n-------------------------------------------------------------->")

	// %d	base 10

	// Declare an integer variable
	var num int = 42

	// Print the integer using %d
	fmt.Printf("Decimal: %d\n", num)

	fmt.Println("\n-------------------------------------------------------------->")

	//%O	base 8 with 0o prefix
	/*
	 the %O format specifier is used to format and print integers in base 8 (octal) with the 0o prefix.
	 Octal notation uses digits from 0 to 7.
	*/
	// Declare an integer variable
	var num2 int = 42

	// Print the integer using %O
	fmt.Printf("Octal: %O\n", num2) //output : Octal: 052

	fmt.Println("\n-------------------------------------------------------------->")

	//%q	a single-quoted character literal safely escaped with Go syntax.

	/*
	 %q format specifier is used to safely escape a single-quoted character literal with Go syntax.
	 It's particularly useful when you need to represent a character literal in a way that is both human-readable
	 and compatible with Go syntax.
	*/
	// Declare a character variable
	var ch byte = 'A'

	// Print the character using %q
	fmt.Printf("Character: %q\n", ch)

	fmt.Println("\n-------------------------------------------------------------->")

	// %x	base 16, with lower-case letters for a-f
	// %X	base 16, with upper-case letters for A-F

	// Declare an integer variable
	var num3 int = 255

	// Print the integer using %x
	fmt.Printf("Lower-case hexadecimal: %x\n", num3)

	// Print the integer using %X
	fmt.Printf("Upper-case hexadecimal: %X\n", num3)

	/*
		output:
		Lower-case hexadecimal: ff
		Upper-case hexadecimal: FF

	*/

	fmt.Println("\n-------------------------------------------------------------->")

	//%U	Unicode format: U+1234; same as "U+%04X"

	// Declare a rune variable
	var r rune = '€'

	// Print the rune and its corresponding code point using %U
	fmt.Printf("Rune: %c\nCode Point: %U\n", r, r)

	/*
	   Rune: €
	   Code Point: U+20AC
	*/

}
