package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func addStudent(scanner *bufio.Scanner, students []*Student) []*Student {
	fmt.Print("Введите фамилию студента: ")
	scanner.Scan()
	lastName := strings.TrimSpace(scanner.Text())

	fmt.Print("Введите имя студента: ")
	scanner.Scan()
	firstName := strings.TrimSpace(scanner.Text())

	fmt.Print("Введите отчество студента: ")
	scanner.Scan()
	middleName := strings.TrimSpace(scanner.Text())

	student := &Student{
		FirstName:  firstName,
		MiddleName: middleName,
		LastName:   lastName,
		Grades:     []int{},
	}

	students = append(students, student)
	fmt.Printf("Студент %s %s %s успешно добавлен!\n", lastName, firstName, middleName)
	return students
}

func addGrade(scanner *bufio.Scanner, students []*Student) []*Student {
	if len(students) == 0 {
		fmt.Println("Нет студентов в журнале!")
		return students
	}

	fmt.Println("\n=== Выберите студента ===")
	for i, student := range students {
		fmt.Printf("%d. %s %s %s\n", i+1, student.LastName, student.FirstName, student.MiddleName)
	}

	fmt.Print("Введите номер студента: ")
	scanner.Scan()
	choiceStr := scanner.Text()

	choice, err := strconv.Atoi(choiceStr)
	if err != nil || choice < 1 || choice > len(students) {
		fmt.Println("Неверный номер студента!")
		return students
	}

	student := students[choice-1]

	fmt.Print("Введите оценку (2-5): ")
	scanner.Scan()
	gradeStr := scanner.Text()

	grade, err := strconv.Atoi(gradeStr)
	if err != nil || grade < 2 || grade > 5 {
		fmt.Println("Неверная оценка! Допустимые значения: 2, 3, 4, 5")
		return students
	}

	student.AddGrade(grade)
	fmt.Printf("Оценка %d добавлена студенту %s %s %s\n",
		grade, student.LastName, student.FirstName, student.MiddleName)
	return students
}

func showAllStudents(students []*Student) {
	if len(students) == 0 {
		fmt.Println("Нет студентов в журнале!")
		return
	}

	fmt.Println("\n=== Журнал студентов ===")
	for i, student := range students {
		fmt.Printf("%d. ", i+1)
		student.DisplayInfo()
	}
}

func deleteStudent(scanner *bufio.Scanner, students []*Student) []*Student {
	if len(students) == 0 {
		fmt.Println("Нет студентов в журнале!")
		return students
	}

	fmt.Println("\n=== Выберите студента для удаления ===")
	for i, student := range students {
		fmt.Printf("%d. %s %s %s\n", i+1, student.LastName, student.FirstName, student.MiddleName)
	}

	fmt.Print("Введите номер студента для удаления: ")
	scanner.Scan()
	choiceStr := scanner.Text()

	choice, err := strconv.Atoi(choiceStr)
	if err != nil || choice < 1 || choice > len(students) {
		fmt.Println("Неверный номер студента!")
		return students
	}

	student := students[choice-1]
	fmt.Printf("Студент %s %s %s удален!\n",
		student.LastName, student.FirstName, student.MiddleName)

	newStudents := make([]*Student, 0, len(students)-1)
	newStudents = append(newStudents, students[:choice-1]...)
	newStudents = append(newStudents, students[choice:]...)

	return newStudents
}
