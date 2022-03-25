package main

import "fmt"

func main() {
	i := 1
	j := &i              // p, of type *int, points to i
	fmt.Println(*j)      // value of variable j = 1
	fmt.Println(j)       // address of variable j
	fmt.Println(&i)      // address of variable i
	fmt.Println(&i == j) // true
	*j = 2               // equivalent to i = 2
	fmt.Println(i)       // 2
	fmt.Println(*j)      // 2

	var x, y int = 0, 1
	fmt.Println(x)  // Print value of variable x
	fmt.Println(&x) // Print address of variable x

	fmt.Println(y)  // Print value of variable y
	fmt.Println(&y) // Print address of variable y

	// The zero value for a pointer of any type is nil
	fmt.Println(&x == &x, &x == &y, &x == nil) // "true false false"
}
