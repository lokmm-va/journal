package main

import "fmt"

type Student struct {
	FirstName  string
	MiddleName string
	LastName   string
	Grades     []int
}

func (s *Student) CalculateAverage() float64 {
	if len(s.Grades) == 0 {
		return 0.0
	}

	sum := 0
	for _, grade := range s.Grades {
		sum += grade
	}
	return float64(sum) / float64(len(s.Grades))
}

func (s *Student) AddGrade(grade int) {
	s.Grades = append(s.Grades, grade)
}

func (s *Student) DisplayInfo() {
	average := s.CalculateAverage()
	fmt.Printf("Студент: %s %s %s\n", s.LastName, s.FirstName, s.MiddleName)
	fmt.Printf("Оценки: %v\n", s.Grades)
	fmt.Printf("Средний балл: %.2f\n\n", average)
}
