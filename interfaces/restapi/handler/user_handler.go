package handler

import (
	"net/http"

	"github.com/finnpn/workout-tracker/interfaces/restapi/models"
	"github.com/finnpn/workout-tracker/pkg/errors"
	"github.com/finnpn/workout-tracker/pkg/ginout"
	"github.com/finnpn/workout-tracker/pkg/log"
	"github.com/finnpn/workout-tracker/usecases/in"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Register(c *gin.Context) {
	g := ginout.NewGinOut(c)
	var register models.Register
	ctx := c.Request.Context()
	if err := c.ShouldBindJSON(&register); err != nil {
		log.Error(errors.ErrParseRequestFail, err)
		g.BadRequest(err.Error())
		return
	}
	input := &in.Register{
		Email:    register.Email,
		Password: register.Password,
	}
	err := h.userUc.Register(ctx, input)
	if err != nil {
		log.Error(errors.ErrRegisterRequestFail, err)
		g.BadRequest(err.Error())
		return
	}

	g.JSONResponse(http.StatusOK, "registered")
}
