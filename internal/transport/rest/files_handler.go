package handler

import (
	"errors"
	"fmt"
	"net/http"
	"net/textproto"

	"github.com/gin-gonic/gin"
	"github.com/salesforceanton/files-portal/internal/logger"
	files_portal "github.com/salesforceanton/files-portal/pkg/domain"
)

const (
	MAX_FILE_SIZE          = 2 * 1024 * 1024
	FORM_DATA_KEY          = "image"
	IMAGE_PNG_CONTENT_TYPE = "image/png"
)

func (h *Handler) GetAll(ctx *gin.Context) {
	// Get user context
	userId, err := h.getUserContext(ctx)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	result, err := h.services.Files.GetFiles(userId)
	if err != nil {
		logger.LogHandlerIssue("get-all", err)
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(
		http.StatusOK, files_portal.FilesResponse{Data: result},
	)
}
func (h *Handler) Upload(ctx *gin.Context) {
	// Get attached file from http request
	request := ctx.Request
	request.Body = http.MaxBytesReader(ctx.Writer, request.Body, MAX_FILE_SIZE)
	file, handler, err := request.FormFile(FORM_DATA_KEY)
	if err != nil {
		logger.LogHandlerIssue("upload", err)
		NewErrorResponse(ctx, http.StatusBadRequest, "Unexpected file size")
		return
	}

	defer file.Close()
	// Check content-type of uploaded file
	err = h.checkContentType(handler.Header, IMAGE_PNG_CONTENT_TYPE)
	if err != nil {
		logger.LogHandlerIssue("upload", err)
		NewErrorResponse(ctx, http.StatusBadRequest, "Unexpected file format")
		return
	}

	// Get user context
	userId, err := h.getUserContext(ctx)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	// Put file into storage and create db record with all file info
	id, err := h.services.AddFileInfo(files_portal.FileItem{
		Filename: handler.Filename,
		Source:   file,
		Size:     handler.Size,
	}, userId)

	if err != nil {
		logger.LogHandlerIssue("create", err)
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, map[string]interface{}{
		"Status": fmt.Sprintf("File record [id]:%d has been created successfully", id),
	})
}

func (h *Handler) checkContentType(header textproto.MIMEHeader, expected string) error {
	ctype := header.Get("Content-Type")
	if ctype != expected {
		return errors.New(fmt.Sprintf("Unexpected content-type shouild be:%s", expected))
	}
	return nil
}
