package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Producer interface {
	Send(string) (string, error)
}

type Handler struct {
	producer Producer
}

func NewHandler(producer Producer) *Handler {
	return &Handler{producer: producer}
}

func (h *Handler) Publish(c *gin.Context) {
	var body struct {
		Message string `json:"message"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.producer.Send(body.Message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "message sent", "response": response})
}
