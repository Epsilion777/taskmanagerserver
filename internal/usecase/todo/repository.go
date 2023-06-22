package todo

import "taskmanagerserver/internal/usecase/interfaces"

type Usecase struct {
	todoStorage interfaces.TodoStorage
}

func NewUsecase(todoStorage interfaces.TodoStorage) *Usecase {
	return &Usecase{
		todoStorage: todoStorage,
	}
}
