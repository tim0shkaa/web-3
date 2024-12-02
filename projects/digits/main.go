package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	fmt.Print("Введите строку: ")
	fmt.Scan(&s)
	var m int = 0

	for _, char := range s {
		digit, err := strconv.Atoi(string(char))
		if err != nil {
			fmt.Println("Ошибка при преобразовании символа в цифру:", err)
			return
		}
		if digit > m {
			m = digit
		}
	}
	fmt.Println(m)
}
