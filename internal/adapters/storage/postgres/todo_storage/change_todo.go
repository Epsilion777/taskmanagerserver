package todostorage

import (
	"context"
	"fmt"
	"taskmanagerserver/internal/domain"
)

const changeTodoQuery = `UPDATE todos SET title = $1, progress = $2, date = $3 WHERE id = $4`

func (p *Storage) ChangeTodo(ctx context.Context, changeTodo *domain.Todo) error {
	res, err := p.db.ExecContext(ctx, changeTodoQuery,
		changeTodo.Title,
		changeTodo.Progress,
		changeTodo.DateEnd,
		changeTodo.ID,
	)
	if err != nil {
		return fmt.Errorf("failed to update todo: %w", err)
	}
	num, err := res.RowsAffected()
	if num == 0 || err != nil {
		return fmt.Errorf("todo not found: %w", err)
	}
	return nil
}
