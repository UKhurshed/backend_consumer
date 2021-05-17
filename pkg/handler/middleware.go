package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userID"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		newResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	if len(headerParts[1]) == 0 {
		newResponse(c, http.StatusUnauthorized, "token is empty")
		return
	}

	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.Set(userCtx, userId)
}

func getUserID(c *gin.Context) (int, error) {
	idFromCtx, ok := c.Get(userCtx)

	if !ok {
		return 0, errors.New("user id not found")
	}

	idInt, ok := idFromCtx.(int)
	if !ok {
		return 0, errors.New("user id is of invalid type")
	}

	return idInt, nil
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

//func corsMiddleware(c *gin.Context){
//	c.Header("Access-Control-Allow-Origin", "*")
//	c.Header("Access-Control-Allow-Methods", "*")
//	c.Header("Access-Control-Allow-Headers", "*")
//	c.Header("Content-Type", "application/json")
//
//	if c.Request.Method != "OPTIONS" {
//		c.Next()
//	} else {
//		c.AbortWithStatus(http.StatusOK)
//	}
//}
