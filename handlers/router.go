package handlers

import (
	"goplay-backend-engineer-test/container"
	"goplay-backend-engineer-test/handlers/file"
	"goplay-backend-engineer-test/handlers/user"
	"goplay-backend-engineer-test/helper"

	"github.com/gin-gonic/gin"
)

type Router struct {
	router    *gin.Engine
	container *container.Container
}

func NewRouter(router *gin.Engine, container *container.Container) *Router {
	return &Router{
		router:    router,
		container: container,
	}
}

func (h *Router) RegisterRouter() {
	h.router.GET("/", func(ctx *gin.Context) {
		helper.WriteSuccess(ctx, "success", gin.H{
			"message": "alive",
		}, nil)
	})

	h.router.GET("/file/:id", file.GetFile(h.container.GetFileUsecase))
	h.router.GET("/file", helper.Authentication(), file.GetListFile(h.container.GetListFileUsecase))
	h.router.POST("/file/upload", helper.Authentication(), file.UploadFile(h.container.UploadFileUsecase))

	h.router.POST("/user/login", user.LoginUser(h.container.UserLoginUsecase))
}
