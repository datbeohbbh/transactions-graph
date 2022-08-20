package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Request struct {
	Action  string `json:"action" binding:"required"`
	Address string `json:"address" binding:"required"`
}

type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func (broker *Broker) routes(c *gin.Context) {
	req := Request{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := broker.HandleRequest(c.Request.Context(), req.Action, req.Address)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": resp})
}
