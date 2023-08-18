package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/salesforceanton/files-portal/internal/logger"
	files_portal "github.com/salesforceanton/files-portal/pkg/domain"
)

func (h *Handler) SignUp(ctx *gin.Context) {
	var request files_portal.User

	if err := ctx.BindJSON(&request); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, "Request is invalid")
		logger.LogHandlerIssue("auth-sign-un", err)
		return
	}

	userId, err := h.services.Authorization.CreateUser(request)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		logger.LogHandlerIssue("auth-sign-un", err)
		return
	}

	ctx.JSON(http.StatusCreated, map[string]interface{}{
		"id": userId,
	})
}

func (h *Handler) SignIn(ctx *gin.Context) {
	var request files_portal.SignInInput

	if err := ctx.BindJSON(&request); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, "Request is invalid")
		logger.LogHandlerIssue("auth-sign-in", err)
		return
	}

	token, err := h.services.Authorization.GenerateAccesssToken(request.Username, request.Password)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		logger.LogHandlerIssue("auth-sign-in", err)
		return
	}

	ctx.JSON(http.StatusCreated, map[string]interface{}{
		"token": token,
	})
}
