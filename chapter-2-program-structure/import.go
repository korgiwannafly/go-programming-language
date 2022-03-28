package main

// import (
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"tempconvert" // package local
// )

// func main() {
// 	for _, arg := range os.Args[1:] {
// 		t, err := strconv.ParseFloat(arg, 64)

// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
// 			os.Exit(1)
// 		}

// 		f := tempconvert.Fahrenheit(t)
// 		c := tempconvert.Celsius(t)
// 		fmt.Printf("%s = %s, %s = %s\n", f, tempconvert.FToC(f), c, tempconvert.CToF(c))
// 	}
// }
