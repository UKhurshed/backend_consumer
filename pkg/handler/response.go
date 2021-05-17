package handler

import (
	"backend_consumer/pkg/logger"
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

type dataResponse struct {
	Data interface{} `json:"data"`
}

type response struct {
	Message string `json:"message"`
}

func newResponse(c *gin.Context, statusCode int, message string) {
	logger.Error(message)
	c.AbortWithStatusJSON(statusCode, response{message})
}
