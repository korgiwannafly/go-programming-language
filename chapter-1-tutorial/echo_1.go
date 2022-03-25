package main

import (
	"os"
)

func main() {
	var s, sep string

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	println(s)
	println(len(os.Args))
}

// Type of loop conditions

// for initialization; condition; post {
// 	zero or more statements
// }

// for condition {
// a traditional "while" loop
// }

// for {
// a traditional infinite loop
// }
