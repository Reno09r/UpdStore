package handler

import (
	"net/http"
	"strconv"

	store "github.com/Reno09r/Store"
	"github.com/gin-gonic/gin"
)


func (h *Handler) аddProduct(c *gin.Context){
	var input store.Product
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Product.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllProductsResponse struct {
	Data []store.Product `json:"data"`
}


func (h *Handler) getAllProducts(c *gin.Context){
	lists, err := h.services.Product.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllProductsResponse{
		Data: lists,
	})
}

func (h *Handler) getProductById(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	list, err := h.services.Product.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}

func (h *Handler) updateProduct(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input store.UpdateProductInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Product.Update(id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) deleteProduct(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.Product.Delete(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) getAllProductsByManufacturer(c *gin.Context) {
    manufacturerID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        newErrorResponse(c, http.StatusBadRequest, "Invalid manufacturer ID")
        return
    }

    products, err := h.services.Product.GetAllByManufacturer(manufacturerID)
    if err != nil {
        newErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }

    c.JSON(http.StatusOK, products)
}

func (h *Handler) getAllProductsByCatalog(c *gin.Context) {
    catalogID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        newErrorResponse(c, http.StatusBadRequest, "Invalid catalog ID")
        return
    }

    products, err := h.services.Product.GetAllByCatalog(catalogID)
    if err != nil {
        newErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }

    c.JSON(http.StatusOK, products)
}