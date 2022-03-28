package main

import "fmt"

func main() {
	p := new(int)   // p, of type *int, points to an unnamed int variable
	fmt.Println(p)  // Print addres of variable p = 0xc0000...
	fmt.Println(*p) // Print value of variable p = 0
	*p = 2
	fmt.Println(p)
	fmt.Println(*p)

	fmt.Println("The New Function")
	i := newInt()
	j := newIntSecond()
	k := new(int)

	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)

	fmt.Printf("%d == %d ? %t\n", i, j, i == j)
	fmt.Printf("%d == %d ? %t\n", i, k, i == k)
}

func newInt() *int {
	return new(int)
}

func newIntSecond() *int {
	var dummy int
	return &dummy
}
