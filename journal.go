package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Journal struct {
	students map[string]Student
}

func NewJournal() *Journal {
	return &Journal{
		students: make(map[string]Student),
	}
}

func containsOnlyLetters(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func (j *Journal) AddStudent() {
	scanner := bufio.NewScanner(os.Stdin)

	var lastName, firstName string

	for {
		fmt.Print("Фамилия: ")
		scanner.Scan()
		lastName = strings.TrimSpace(scanner.Text())
		if containsOnlyLetters(lastName) && lastName != "" {
			break
		}
		fmt.Println("Ошибка! Фамилия должна содержать только буквы")
	}

	for {
		fmt.Print("Имя: ")
		scanner.Scan()
		firstName = strings.TrimSpace(scanner.Text())
		if containsOnlyLetters(firstName) && firstName != "" {
			break
		}
		fmt.Println("Ошибка! Имя должно содержать только буквы")
	}

	fullName := lastName + " " + firstName

	if _, exists := j.students[fullName]; exists {
		fmt.Println("Студент уже существует!")
		return
	}

	grades := j.inputGrades()

	j.students[fullName] = Student{
		FirstName: firstName,
		LastName:  lastName,
		Grades:    grades,
	}

	fmt.Println("Студент добавлен!")
}

func (j *Journal) inputGrades() []int {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Оценки через пробел (2-5): ")
		scanner.Scan()
		gradesInput := strings.Fields(scanner.Text())

		grades := make([]int, 0, len(gradesInput))
		valid := true

		for _, gradeStr := range gradesInput {
			grade, err := strconv.Atoi(gradeStr)
			if err != nil || grade < 2 || grade > 5 {
				fmt.Println("Ошибка! Оценки должны быть числами от 2 до 5")
				valid = false
				break
			}
			grades = append(grades, grade)
		}

		if valid {
			return grades
		}
	}
}

func (j *Journal) ModifyStudentGrades() {
	if len(j.students) == 0 {
		fmt.Println("Студентов нет")
		return
	}

	j.ShowAllStudents()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите полное имя студента (Фамилия Имя): ")
	scanner.Scan()
	fullName := strings.TrimSpace(scanner.Text())

	student, exists := j.students[fullName]
	if !exists {
		fmt.Println("Студент не найден!")
		return
	}

	fmt.Printf("Текущие оценки студента %s: ", fullName)
	for i, grade := range student.Grades {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(grade)
	}
	fmt.Println()

	newGrades := j.inputGrades()
	student.Grades = newGrades
	j.students[fullName] = student

	fmt.Println("Оценки обновлены!")
}

func (j *Journal) FilterByAverage(min, max float64) {
	fmt.Printf("Студенты со средним баллом от %.1f до %.1f:\n", min, max)
	found := false

	for _, student := range j.students {
		average := student.AverageGrade()
		if average >= min && average <= max {
			student.DisplayInfo()
			found = true
		}
	}

	if !found {
		fmt.Println("Студенты не найдены")
	}
}

func (j *Journal) ShowAllStudents() {
	if len(j.students) == 0 {
		fmt.Println("Студентов нет")
		return
	}

	fmt.Println("Все студенты:")
	for _, student := range j.students {
		student.DisplayInfo()
	}
}

func (j *Journal) ShowStatistics() {
	if len(j.students) == 0 {
		fmt.Println("Студентов нет")
		return
	}

	totalAverage := 0.0
	count := 0

	for _, student := range j.students {
		if len(student.Grades) > 0 {
			totalAverage += student.AverageGrade()
			count++
		}
	}

	if count > 0 {
		fmt.Printf("Общий средний балл группы: %.2f\n", totalAverage/float64(count))
	}
}
