package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/salesforceanton/files-portal/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("auth")
	{
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/sign-in", h.SignIn)
	}
	api := router.Group("api", h.userIdentity)
	{
		files := api.Group("files")
		{
			files.GET("/", h.GetAll)
			files.POST("/", h.Create)
		}
	}

	return router
}

func (h *Handler) getUrlParam(ctx *gin.Context, param string) (int, error) {
	result, err := strconv.Atoi(ctx.Param(param))
	return result, err
}
