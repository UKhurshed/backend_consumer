package handler

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Error struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string){
	c.AbortWithStatusJSON(statusCode, Error{message})
	log.Fatal(message)
}

type statusResponse struct {
	Status string `json:"status"`
}

type DataResponse struct {
	Data interface{} `json:"data"`
}
