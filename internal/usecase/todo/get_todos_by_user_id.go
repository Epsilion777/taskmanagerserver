package todo

import (
	"context"
	"taskmanagerserver/internal/domain"
)

func (uc *Usecase) GetTodosByUserID(ctx context.Context, userID string) ([]*domain.Todo, error) {
	return uc.todoStorage.GetTodosByUserID(ctx, userID)
}
