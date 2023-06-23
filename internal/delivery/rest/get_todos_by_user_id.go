package rest

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type TodoJson struct {
	ID       int       `json:"id"`
	UserID   string    `json:"user_id"`
	Title    string    `json:"title"`
	Progress int       `json:"progress"`
	DateEnd  time.Time `json:"date"`
}

type GetTodosResponse struct {
	Todos []TodoJson `json:"todos"`
}

func (h *Handler) GetTodosByUserID(c *gin.Context) {
	userID := c.Param("id")

	todos, err := h.todoUseCase.GetTodosByUserID(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	todosJSON := make([]TodoJson, 0, len(todos))

	for _, v := range todos {
		newTodo := TodoJson{
			ID:       v.ID,
			UserID:   v.UserID,
			Title:    v.Title,
			Progress: v.Progress,
			DateEnd:  v.DateEnd,
		}
		todosJSON = append(todosJSON, newTodo)
	}

	c.JSON(http.StatusOK, GetTodosResponse{
		Todos: todosJSON,
	})
}
