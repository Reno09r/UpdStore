package handler

import (
	"net/http"

	store "github.com/Reno09r/Store"
	"github.com/gin-gonic/gin"
)

type getAllBoughtProductsResponse struct{
	Data []store.BoughtProducts `json:"data"`
}

func (h *Handler) ConfirmBuy(c *gin.Context){
	var input store.UserCardInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	err = h.services.Buy.Confirm(input, userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) BuyedProducts(c *gin.Context){
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	lists, err := h.services.Buy.BoughtProducts(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllBoughtProductsResponse{
		Data: lists,
	})
}