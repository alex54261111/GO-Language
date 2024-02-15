package main

import (
	"errors"
	"fmt"
)

func validate(input int) error {
	if input < 0 {
		return errors.New("輸入小於零")
	} else if input > 100 {
		return errors.New("不可大於100")
	} else if input%7 == 0 {
		return errors.New("不可為7的倍數")
	} else {
		return nil
	}
}

func main() {
	input := 21
	if err := validate(input); err != nil {//起始賦值敘述
		fmt.Print(err)
	} else if input%2 == 0 {
		fmt.Println(input, "偶數")
	} else {
		fmt.Println(input, "奇數")
	}
}
