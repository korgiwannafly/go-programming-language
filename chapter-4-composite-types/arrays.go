package main

import "fmt"

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

// go’s arrays are values
func main() {
	var a = [...]int{1: 2, 3: 4} // array of 3 integers

	// print the indices and elements.
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	// to avoid the overhead of copying array using pointer
	var b = &a // b is a pointer to array a
	fmt.Println(a[0], a[1])
	fmt.Println(b[0], b[1])
	for index, value := range b {
		b[index] += 1
		fmt.Println(index, value)
	}

	for index, value := range a {
		fmt.Println(index, value)
	}

	fmt.Println(a[0])        // print the first element
	fmt.Println(a[len(a)-1]) // print the last element, a[2]

	// print the indices and elements.
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	// Print the elements only.
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	fmt.Println(r[2]) // 0

	q = [...]int{1, 2, 3}
	fmt.Printf("%T\n", q) // [3]int

	symbol := [...]string{USD: "$", EUR: "9", GBP: "!", RMB: ""}
	fmt.Println(EUR, symbol[EUR]) // 1 9
}
