package main

import (
	"fmt"
	"math"
)

func T(k, p, v float64) float64 {
	return 6 / W(k, p, v)
}

func W(k, p, v float64) float64 {
	return math.Sqrt(k / M(p, v))
}

func M(p, v float64) float64 {
	return p * v
}

func main() {
	var k, p, v float64
	fmt.Println("Введите через пробел значение переменных k, p, v: ")

	// Проверка ввода
	if _, err := fmt.Scan(&k, &p, &v); err != nil {
		fmt.Println("Ошибка: Введите корректные числовые значения.")
		return
	}

	fmt.Println(T(k, p, v))
}
