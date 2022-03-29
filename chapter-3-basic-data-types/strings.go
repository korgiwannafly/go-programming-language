package main

import "fmt"

func main() {
	s := "hello, world"
	fmt.Println(len(s))     // "12"
	fmt.Println(s[0], s[7]) // "104 119"  ('h' and 'w')

	// c := s[len(s)]
	// fmt.Println(c) // panic: index out of range

	fmt.Println(s[0:5])            // "hello"
	fmt.Println(s[:5])             // "hello"
	fmt.Println(s[7:])             // "world"
	fmt.Println(s[:])              // "hello, world"
	fmt.Println("goodbye" + s[5:]) // "goodbye, world"

	i := "left foot"
	j := i
	i += ", right foot"
	fmt.Println(i) // "left foot, right foot"
	fmt.Println(j) // "left foot"
}
