package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/salesforceanton/files-portal/internal/logger"
)

const (
	AUTH_HEADER = "Authorization"
	USER_CTX    = "user_id"
)

func (h *Handler) userIdentity(ctx *gin.Context) {
	authHeader := ctx.GetHeader(AUTH_HEADER)
	if authHeader == "" {
		logger.LogHandlerIssue("user-identity", errors.New("Authorization Header is empty"))
		NewErrorResponse(ctx, http.StatusUnauthorized, "Authorization Header is empty")
		return
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		logger.LogHandlerIssue("user-identity", errors.New("Authorization Header is invalid"))
		NewErrorResponse(ctx, http.StatusUnauthorized, "Authorization Header is invalid")
		return
	}

	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		logger.LogHandlerIssue("user-identity", errors.New(fmt.Sprintf("Access Token is invalid: %s", err.Error())))
		NewErrorResponse(ctx, http.StatusInternalServerError, fmt.Sprintf("Access Token is invalid: %s", err.Error()))
		return
	}

	ctx.Set(USER_CTX, userId)
}

func (h *Handler) getUserContext(ctx *gin.Context) (int, error) {
	userId, ok := ctx.Get(USER_CTX)
	if !ok {
		logger.LogHandlerIssue("api", errors.New("User id is not found"))
		NewErrorResponse(ctx, http.StatusInternalServerError, "User id is not found")
		return 0, errors.New("User id is not found")
	}

	return userId.(int), nil
}
