package todo

import (
	"context"
)

func (uc *Usecase) DeleteTodo(ctx context.Context, id string) error {
	return uc.todoStorage.DeleteTodo(ctx, id)
}
