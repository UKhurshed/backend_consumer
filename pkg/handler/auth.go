package handler

import (
	"backend_consumer/pkg/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type signInInput struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary SignIn
// @Tags auth
// @Description login
// @ID login
// @Accept  json
// @Produce  json
// @Param input body signInInput true "credentials"
// @Success 200 {string} string "token"
// @Failure 400,404 {object} Error
// @Failure 500 {object} Error
// @Failure default {object} Error
// @Router /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context){
	var input signInInput
	if err:= c.BindJSON(&input); err != nil{
		newResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Email, input.Password)

	if err != nil{
		newResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

// @Summary SignUp
// @Tags auth
// @Description create account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body domain.User true "account info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} Error
// @Failure 500 {object} Error
// @Failure default {object} Error
// @Router /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context){
	var user domain.User
	if err := c.BindJSON(&user); err != nil{
		newResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.services.Authorization.CreateUser(user)

	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}
