package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Request struct {
	Class  string `json:"class" binding:"required"`
	Action string `json:"action" binding:"required"`
	Data   any    `json:"data"`
}

type Response struct {
	Error   bool   `json:"error"`
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func createResponse(Error_ bool, Status_, Message_ string, Data_ any) *Response {
	return &Response{
		Error:   Error_,
		Status:  Status_,
		Message: Message_,
		Data:    Data_,
	}
}

func (broker *Broker) routes(c *gin.Context) {
	req := Request{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp := broker.HandleRequest(c.Request.Context(), &req)
	c.JSON(http.StatusOK, gin.H{"response": resp})
}
