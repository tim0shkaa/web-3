package main

import (
	"fmt"
)

func main() {
	var input, result string
	fmt.Print("Введите строку: ")
	if _, err := fmt.Scan(&input); err != nil {
		fmt.Println("Ошибка ввода, используйте только буквы английского алфавита")
		return
	}
	for _, char := range input {
		result += string(char)
		result += "*"
	}
	if len(result) > 0 {
		result = result[:len(result)-1]
	}
	fmt.Println(result)
}
