package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
