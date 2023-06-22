package interfaces

import (
	"context"
	"taskmanagerserver/internal/domain"
)

type TodoUseCase interface {
	GetTodosByUserID(ctx context.Context, userID string) ([]*domain.Todo, error)
	CreateTodo(ctx context.Context, createTodo *domain.Todo) error
	ChangeTodo(ctx context.Context, changeTodo *domain.Todo) error
	DeleteTodo(ctx context.Context, id string) error
}
