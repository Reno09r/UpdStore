package handler

import (
	"net/http"

	store "github.com/Reno09r/Store"
	"github.com/gin-gonic/gin"
)



func (h *Handler) signUp(c *gin.Context){
	var input store.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.services.Authentication.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}



type signInInput struct{
	Username string `json:"username" bilding:"required"`
	Password string `json:"password" bilding:"required"`
}

func (h *Handler) signIn(c *gin.Context){
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authentication.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

