package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

const (
	AUTH_HEADER = "Authorization"
	USER_CTX    = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(AUTH_HEADER)

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

	c.Set(USER_CTX, userId)
}

func getUserIdFromContext(c *gin.Context) (uint64, error) {
	const ERROR_MESSAGE = "userid is not found in context"
	userId, ok := c.Get(USER_CTX)
	if !ok {
		newErrorResponse(c, http.StatusBadRequest, ERROR_MESSAGE)
		return 0, fmt.Errorf(ERROR_MESSAGE)
	}
	return userId.(uint64), nil
}

func getIdFromParam(c *gin.Context) (uint64, error) {
	const ERROR_MESSAGE = "id is not provided"
	listIdAsStr := c.Param("id")
	if len(listIdAsStr) == 0 {
		newErrorResponse(c, http.StatusBadRequest, ERROR_MESSAGE)
		return 0, fmt.Errorf(ERROR_MESSAGE)
	}
	listId, err := strconv.Atoi(listIdAsStr)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return 0, err
	}

	return uint64(listId), nil

}
