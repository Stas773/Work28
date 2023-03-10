package storage

import "D/Works/GO/27work/pkg/student"

type Storage interface {
	Get(string) (*student.Student, error)
	Put(*student.Student) (int, error)
	PrintStorage() map[string]*student.Student
}
