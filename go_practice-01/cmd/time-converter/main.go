package main

import "fmt"

func main() {
	var minutes int
	fmt.Print("Введите количество минут: ")
	_, err := fmt.Scan(&minutes)
	if err == nil {
		if minutes < 0 {
			fmt.Println("отрицательные значения не допустимы")
		} else {
			hours := minutes / 60
			miniminutes := minutes % 60
			fmt.Printf("Полных часов: %d \n", hours)
			fmt.Printf("Оставшихся минут: %d \n", miniminutes)
		}
	} else {
		fmt.Println("Ошибка чтения числа")
	}
}
