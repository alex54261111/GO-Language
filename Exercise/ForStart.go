package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		for a := 5; a > i; a-- {
			fmt.Print(" ")
		}
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		for a := 0; a < i; a++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
	for a := 0; a < 10; a++ {
		fmt.Print("_")
	}
}
