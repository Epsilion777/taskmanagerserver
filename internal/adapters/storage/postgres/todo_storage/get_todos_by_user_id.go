package todostorage

import (
	"context"
	"fmt"
	"taskmanagerserver/internal/domain"
)

const getTodosQuery = `SELECT * FROM todos WHERE userid = $1;`

func (p *Storage) GetTodosByUserID(ctx context.Context, userID string) ([]*domain.Todo, error) {
	todos := make([]*domain.Todo, 0, 0)

	rows, err := p.db.QueryContext(ctx, getTodosQuery, userID)
	if err != nil {
		return nil, fmt.Errorf("cannot get todos: %w", err)
	}
	for rows.Next() {
		todo := &domain.Todo{}
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Progress, &todo.DateEnd, &todo.UserID); err != nil {
			return nil, fmt.Errorf("cannot parse todo: %w", err)
		}
		todos = append(todos, todo)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return todos, nil
}
