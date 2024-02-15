package main

import (
	"fmt"
	"time"
)

var a string = "str" //
var (
	a1 string    = "str"
	b1 bool      = false
	c1 time.Time = time.Now()
)
var (
	a2 = "str"
	b2 = false
	c2 = time.Now()
)

func main() {
	a3 := "str" //:= 短數宣告，只能在函式宣告
	b3 := false
	c3 := time.Now()

	a4, b4, c4 := "str", true, time.Now()
	a5, b5, c5 := getConfig()

	fmt.Println(a1, b1, c1)
	fmt.Println(a2, b2, c2)
	fmt.Println(a3, b3, c3)
	fmt.Println(a4, b4, c4)
	fmt.Println(a5, b5, c5)
}

func getConfig() (bool, string, time.Time) { //函式回傳多重變數
	return true, "str", time.Now()
}
