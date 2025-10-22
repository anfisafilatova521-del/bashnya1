package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Scan(&number)

	for number < 12307 {
		fmt.Println("Текущее число:", number)// ← отладка
		if number < 0 {
			number = number * (-1)
		} else if number%7 == 0 {
			number = number * 39
		} else if number%9 == 0 {
			number = number*13 + 1
			continue
		} else {
			number = (number + 2) * 3
		}
fmt.Println("Текущее число после выполнения первого блока :", number)// ← отладка
		if number%9 == 0 && number%13 == 0 {
			fmt.Println("service error ")
			return
		} else {
			number = number + 1
		}

	}

	fmt.Printf("Цикл завершён: итоговое значение числа — %d\n,число>12307", number)
}
