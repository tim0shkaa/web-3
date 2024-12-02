package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	fmt.Print("Введите целое число: ")
	fmt.Scan(&input)

	result := ""

	for _, char := range input {
		digit, err := strconv.Atoi(string(char))
		if err != nil {
			fmt.Println("Ошибка при преобразовании символа в цифру:", err)
			return
		}
		square := digit * digit
		result += strconv.Itoa(square)
	}

	fmt.Println("Результат:", result)
}
