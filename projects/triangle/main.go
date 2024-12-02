package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64
	fmt.Println("Введите два положительных числа через пробел:")

	_, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	if a <= 0 || b <= 0 {
		fmt.Println("Оба числа должны быть положительными")
		return
	}

	fmt.Println(math.Sqrt(a*a + b*b))
}
