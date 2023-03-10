package memstorage

import (
	"D/Works/GO/27work/pkg/student"
	"errors"
)

type MemStorage struct {
	Students map[string]*student.Student
}

func NewMemStore() *MemStorage {
	return &MemStorage{
		Students: make(map[string]*student.Student, 0),
	}
}

func (ms *MemStorage) Put(value *student.Student) (int, error) {
	ms.Students[value.Name] = value
	if ms.Students == nil {
		return -1, errors.New("ошибка добавления данных в хранилище")
	} else {
		return 0, nil
	}
}
func (ms *MemStorage) Get(name string) (*student.Student, error) {
	if ms.Students[name] == nil {
		return nil, errors.New("студент не найден в хранилище")
	} else {
		return ms.Students[name], nil
	}
}

func (ms *MemStorage) PrintStorage() map[string]*student.Student {
	return ms.Students
}
