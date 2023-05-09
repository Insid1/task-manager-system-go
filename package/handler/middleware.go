package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authHeader = "Authorization"
	userCtx    = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authHeader)

	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")

	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	userId, err := h.services.Authorization.ParseToken(headerParts[1])

	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
	}

	c.Set(userCtx, userId)
}

func getUserIdFromContext(c *gin.Context) (uint64, error) {
	const ERROR_MESSAGE = "userid is not found in context"
	userId, ok := c.Get(userCtx)

	if !ok {
		return 0, fmt.Errorf(ERROR_MESSAGE)
	}

	return userId.(uint64), nil
}

func getIdFromParam(c *gin.Context) (uint64, error) {
	const ERROR_MESSAGE = "id is not provided"
	listIdAsStr := c.Param("id")

	if len(listIdAsStr) == 0 {
		return 0, fmt.Errorf(ERROR_MESSAGE)
	}
	listId, err := strconv.ParseUint(listIdAsStr, 10, 64)

	if err != nil {
		return 0, err
	}

	return listId, nil

}

func getItemIdFromParam(c *gin.Context) (uint64, error) {
	const ERROR_MESSAGE = "invalid item id"

	ItemIdAsStr := c.Param("item_id")

	if len(ItemIdAsStr) == 0 {
		return 0, fmt.Errorf(ERROR_MESSAGE)
	}

	itemId, err := strconv.ParseUint(ItemIdAsStr, 10, 64)

	if err != nil {
		return 0, err
	}

	return itemId, nil
}
