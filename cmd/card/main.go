package main

import "fmt"

func main() {
	var name string = "Николай"
	var group string = "Учебная-01"
	course := "Бэкенд-разработка на языке Go"
	finished := 1
	interesting := true
	fmt.Println("Участник:", name)
	fmt.Println("Группа:", group)
	fmt.Println("Курс:", course)
	fmt.Println("Завершено работ:", finished)
	fmt.Println("Интерес к бэкенду:", interesting)
}
