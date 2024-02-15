package main

import "fmt"

const num int = 90

const (
	a1 int  = 100
	a2 bool = false
)

const (
	sunday = iota
	monday
	tuesday
	wednesday
	thursday
	friday
	saturday
)

func main() {
	fmt.Print(num, a1, a2, sunday, monday, tuesday, wednesday, thursday, friday, saturday)
}
