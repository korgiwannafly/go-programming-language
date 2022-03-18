package main

import (
	"strings"
)

func main() {
	str := []string{"Go", "Programming", "Language"}

	for _, arg := range []int{1, 2, 3} {
		println(arg)
	}

	println(strings.Join(str, " "))
}

// _ (that is, an underscore) using when not using index of for loop
