package main

import "fmt"

func main() {
	a := [10]int{0, 1, 2, 3, 4, 5}
	fmt.Println("original ", a[:])    // [5 4 3 2 1 0]
	fmt.Println("before reverses", a) // [0 1 2 3 4 5]
	reverse(a[:])
	fmt.Println("after reverses", a) // [5 4 3 2 1 0]
	fmt.Println("length: ", len(a))
	fmt.Println("capacity: ", cap(a))

	s := []int{0, 1, 2, 3, 4, 5}
	// rotate s left by two positions.
	reverse(s[:2])
	reverse(s[2:])
	reverse(s)
	fmt.Println(s) // "[2 3 4 5 0 1]"

	x := [...]string{"x", "x"}
	y := [...]string{"y", "y"}
	fmt.Println(equal(x[:], y[:]))
}

// reverse reverses a slice of ints in place.
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}

	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}

	return true
}
