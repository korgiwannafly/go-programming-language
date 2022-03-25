package main

import (
	"fmt"
)

func main() {
	const freezingF, boilingF = 32.0, 212.0
	println("%g°F = %g°C", freezingF, fToC(freezingF)) // %g°F = %g°C +3.200000e+001 +0.000000e+000
	println("%g°F = %g°C", boilingF, fToC(boilingF))   // %g°F = %g°C +2.120000e+002 +1.000000e+002

	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF)) // "32°F = 0°C"
	fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF))   // "212°F = 100°C"
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
