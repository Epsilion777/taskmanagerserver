package rest

import (
	"taskmanagerserver/internal/usecase/interfaces"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	todoUseCase interfaces.TodoUseCase
}

func NewHandler(todoUseCase interfaces.TodoUseCase, router *gin.Engine) *Handler {
	h := &Handler{
		todoUseCase: todoUseCase,
	}

	router.GET("/api/todos/:id", h.GetTodosByUserID)
	router.POST("/api/createtodo", h.CreateTodo)
	router.DELETE("/api/deletetodo/:id", h.DeleteTodo)
	router.PUT("/api/changetodo", h.ChangeTodo)
	return h
}
