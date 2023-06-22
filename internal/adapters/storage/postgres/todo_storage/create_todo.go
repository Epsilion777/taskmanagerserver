package todostorage

import (
	"context"
	"fmt"
	"taskmanagerserver/internal/domain"
)

const createTodoQuery = `INSERT INTO todos (title, progress, date, userid) values ($1, $2, $3, $4);`

func (p *Storage) CreateTodo(ctx context.Context, createTodo *domain.Todo) error {
	_, err := p.db.ExecContext(ctx, createTodoQuery,
		createTodo.Title,
		createTodo.Progress,
		createTodo.DateEnd,
		createTodo.UserID,
	)
	if err != nil {
		return fmt.Errorf("failed to insert todo: %w", err)
	}
	return nil
}
