package main

import "fmt"

var p = f()

func f() *int {
	v := 0
	return &v
}

func main() {
	fmt.Println(p)
	fmt.Println(p == p)
	fmt.Println(f() == f())

	v := 0
	incr(&v) // side effect: v is now 1
	incr(&v) // 2 (and v is 2)
}

func incr(p *int) int {
	fmt.Println("p = ", p)
	fmt.Println("*p = ", *p)

	*p++ // increments what p points to; does not change p

	fmt.Println("p = ", p)   // print address of variable p
	fmt.Println("*p = ", *p) // *p is value of address p
	return *p
}
