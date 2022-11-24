package handler

import (
	usecases "github.com/finnpn/workout-tracker/usecases/user"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	userUc *usecases.AuthUserUc
}

func NewHandler(userUc *usecases.AuthUserUc) *Handler {
	return &Handler{
		userUc: userUc,
	}
}

//TODO define group of router

func (h *Handler) ConfigureNotAuth(router *gin.RouterGroup) {
	router.POST("/register", h.Register)
	//router.POST("/login", h.Login)
}
