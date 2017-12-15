package main

import "fmt"

func main() {
	x := 1
	p := &x         // p, of type *int, points to x
	fmt.Println(*p) // "1"
	*p = 2          // equivalent to x = 2
	fmt.Println(x)
	fmt.Printf("%T \n", p)  // of type *int (pointer to int)
	fmt.Printf("%T \n", *p) // of type int
	fmt.Printf("%T \n", x)  // of type int
}
