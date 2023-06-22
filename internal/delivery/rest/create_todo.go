package rest

import (
	"net/http"
	"strconv"
	"taskmanagerserver/internal/domain"
	"time"

	"github.com/gin-gonic/gin"
)

type CreateTodoBody struct {
	Title    string `json:"title"`
	Progress string `json:"progress"`
	DateEnd  string `json:"date"`
	User     string `json:"user"`
}

func (h *Handler) CreateTodo(c *gin.Context) {
	var requestBody CreateTodoBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	progressString, err := strconv.Atoi(requestBody.Progress)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	DateEnd, err := time.Parse("2006-01-02", requestBody.DateEnd)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.todoUseCase.CreateTodo(c.Request.Context(), &domain.Todo{
		Title:    requestBody.Title,
		Progress: progressString,
		DateEnd:  DateEnd,
		UserID:   requestBody.User,
	}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "todo created"})
}
