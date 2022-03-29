package main

import "fmt"

type Weekday int

type Flags uint

const (
	FlagUp           Flags = 1 << iota // is up
	FlagBroadcast                      // supports broadcast access capability
	FlagLoopback                       // is a loopback interface
	FlagPointToPoint                   // belongs to a point-to-point link
	FlagMulticast                      // supports multicast access capability
)

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

const (
	a = 1
	b
	c = 2
	d
)

func main() {
	fmt.Println(a, b, c, d)                                                           // 1 1 2 2
	fmt.Println(Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday)       // 1 2 3 4 5 6 0
	fmt.Println(FlagUp, FlagBroadcast, FlagLoopback, FlagPointToPoint, FlagMulticast) // 1 2 4 8 16
}
