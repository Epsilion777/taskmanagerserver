package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) DeleteTodo(c *gin.Context) {
	id := c.Param("id")

	err := h.todoUseCase.DeleteTodo(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "todo deleted"})
}
