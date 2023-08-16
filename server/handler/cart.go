package handler

import (
	"net/http"
	"strconv"

	store "github.com/Reno09r/Store"
	"github.com/gin-gonic/gin"
)

func (h *Handler) insertProductInCart(c *gin.Context){
	var input store.Cart
	customerId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Cart.Insert(input, customerId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
        "id": id,
    })
}

type getAllProductsInCartResponse struct {
	Data []store.Cart `json:"data"`
}

func (h *Handler) getProductsFromCart(c *gin.Context){
	customerId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	lists, err := h.services.Cart.Get(customerId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllProductsInCartResponse{
		Data: lists,
	})
}

func (h *Handler) deleteProductFromCart(c *gin.Context){
	customerId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.Cart.Delete(id, customerId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
