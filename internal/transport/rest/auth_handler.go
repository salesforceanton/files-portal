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
		logger.LogHandlerIssue("auth", err)
		return
	}

}

func (h *Handler) SignIn(ctx *gin.Context) {}
