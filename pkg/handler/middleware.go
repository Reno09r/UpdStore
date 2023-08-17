package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx = "userId"
)

func (h *Handler)identity(isAdminRequired bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader(authorizationHeader)
		if header == "" {
			newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
			c.Abort()
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 {
			newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
			c.Abort()
			return
		}

		ID, err := h.services.ParseToken(headerParts[1])
		if err != nil {
			newErrorResponse(c, http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		if isAdminRequired {
			isAdmin, err := h.services.Authorization.CurrentUserIsAdmin(ID)
			if err != nil {
				newErrorResponse(c, http.StatusInternalServerError, err.Error())
				c.Abort()
				return
			}

			if !isAdmin {
				newErrorResponse(c, http.StatusUnauthorized, "Only admins are allowed")
				c.Abort()
				return
			}
		} else 
		{
			isCustomer, err := h.services.Authorization.CurrentUserIsCustomer(ID)
			if err != nil {
				newErrorResponse(c, http.StatusInternalServerError, err.Error())
				c.Abort()
				return
			}

			if !isCustomer {
				newErrorResponse(c, http.StatusUnauthorized, "Only customers are allowed")
				c.Abort()
				return
			}
		}

		c.Set(userCtx, ID)
	}
}
func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("user id is of invalid type")
	}

	return idInt, nil
}