package todostorage

import (
	"context"
	"fmt"
)

const deleteTodoQuery = `DELETE FROM todos WHERE id = $1`

func (p *Storage) DeleteTodo(ctx context.Context, id string) error {
	res, err := p.db.ExecContext(ctx, deleteTodoQuery, id)
	if err != nil {
		return fmt.Errorf("cannot delete todo: %w", err)
	}
	num, err := res.RowsAffected()
	if num == 0 || err != nil {
		return fmt.Errorf("todo not found: %w", err)
	}
	return nil
}
