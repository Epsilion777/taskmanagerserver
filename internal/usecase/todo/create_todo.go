package todo

import (
	"context"
	"taskmanagerserver/internal/domain"
)

func (uc *Usecase) CreateTodo(ctx context.Context, createTodo *domain.Todo) error {
	return uc.todoStorage.CreateTodo(ctx, createTodo)
}
