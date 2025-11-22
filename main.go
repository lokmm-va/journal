package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	students := make([]*Student, 0)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n Журнал оценок")
		fmt.Println("1. Посмотреть журнал")
		fmt.Println("2. Добавить студента")
		fmt.Println("3. Добавить оценку студенту")
		fmt.Println("4. Удалить студента")
		fmt.Println("5. Выйти")
		fmt.Print("Выберите действие: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			showAllStudents(students)
		case "2":
			students = addStudent(scanner, students)
		case "3":
			students = addGrade(scanner, students)
		case "4":
			students = deleteStudent(scanner, students)
		case "5":
			fmt.Println("Выход из программы.")
			return
		default:
			fmt.Println("Пожалуйста, выберите действие из списка.")
		}
	}
}
