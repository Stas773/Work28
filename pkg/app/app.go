package app

import (
	"D/Works/GO/27work/pkg/storage"
	"D/Works/GO/27work/pkg/student"
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type App struct {
	Repository storage.Storage
}

func (a *App) Run() {
	for {
		if input, ok := a.InputNextStudent(); ok {
			a.StoreStudent(input)
		} else {
			a.PrintStorage()
			break
		}
	}
}

func (a *App) InputNextStudent() (student.Student, bool) {
	newStudent := student.Student{}
	fmt.Println("Введите имя студента, возраст и курс через пробел:")
	for {
		sc := bufio.NewReader(os.Stdin)
		reader, err := sc.ReadString('\n')
		if err == io.EOF {
			fmt.Println("_________________________")
			fmt.Println("Весь список студентов:")
			return newStudent, false
		}
		readerFields := strings.Fields(reader)
		if len(readerFields) < 3 {
			fmt.Println("Необходимо ввести имя, возраст и курс, попробуйте снова:")
			continue
		}

		studentName := readerFields[0]
		studentAge, errAge := strconv.Atoi(readerFields[1])
		studentGrade, errGrade := strconv.Atoi(readerFields[2])

		_, errName := strconv.Atoi(studentName)
		if errName == nil {
			fmt.Println("Неверно указано имя, попробуйте снова:")
			continue
		}

		if errAge != nil || errGrade != nil {
			fmt.Println("Неверно указан возраст или курс, попробуйте снова:")
			continue
		}
		newStudent = student.Student{
			Name:  studentName,
			Age:   studentAge,
			Grade: studentGrade,
		}
		return newStudent, true
	}
}

func (a *App) StoreStudent(input student.Student) {
	if _, err := a.Repository.Get(input.Name); err != nil {
		a.Repository.Put(&input)
	} else {
		fmt.Println("Ошибка добавления в список")
	}
}

func (a *App) PrintStorage() {
	for _, v := range a.Repository.PrintStorage() {
		fmt.Println(v.Name, v.Age, v.Grade)
	}
}
