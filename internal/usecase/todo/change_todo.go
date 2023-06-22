package todo

import (
	"context"
	"taskmanagerserver/internal/domain"
)

func (uc *Usecase) ChangeTodo(ctx context.Context, changeTodo *domain.Todo) error {
	return uc.todoStorage.ChangeTodo(ctx, changeTodo)
}
