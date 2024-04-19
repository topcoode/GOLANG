package main

import "fmt"

//*Point(p)        // same as *(Point(p))

type Point struct {
	x, y int
}

func main() {
	p := &Point{3, 4}
	//fmt.Println(*Point(p)) // Output: {3 4}
	fmt.Println(p)
}

//(*Point)(p)      // p is converted to *Point
/*
	type Point struct {
		x, y int
	}

	func main() {
		p := Point{3, 4}
		fmt.Println((*Point)(&p)) // Output: &{3 4}
}
*/

// <-chan int(c)    // same as <-(chan int(c))
/*
	func main() {
    c := make(chan int)
    c <- 42
    fmt.Println(<-chan int(c)) // Output: 42
}
*/

// (<-chan int)(c)  // c is converted to <-chan int
/*
func main() {
    c := make(chan int)
    c <- 42
    fmt.Println((<-chan int)(c)) // Output: 42
}
*/
// func()(x)        // function signature func() x
// (func())(x)      // x is converted to func()
// (func() int)(x)  // x is converted to func() int
// func() int(x)    // x is converted to func() int (unambiguous)
