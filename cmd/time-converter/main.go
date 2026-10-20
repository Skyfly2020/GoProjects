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
			fmt.Printf("Полных часов: %d \n", minutes/60)
			fmt.Printf("Оставшихся минут: %d \n", minutes%60)
		}
	} else {
		fmt.Println("Ошибка чтения числа")
	}
}
