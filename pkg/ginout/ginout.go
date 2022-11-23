package ginout

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GinOut struct {
	ginCtx *gin.Context
}

func NewGinOut(c *gin.Context) *GinOut {
	return &GinOut{
		ginCtx: c,
	}
}

// JSONResponse ...
func (g *GinOut) JSONResponse(code int, data interface{}) {
	g.ginCtx.JSON(code, gin.H{
		"data": data,
	})
}

// BadRequest ...
func (g *GinOut) BadRequest(message string) {
	g.JSONError(http.StatusBadRequest, "bad_request", message)
	g.ginCtx.Abort()
}

// JSONError response json format error
// this will Abort other handlers
func (g *GinOut) JSONError(statusCode int, code string, message string) {
	errorItem := gin.H{
		"code":    code,
		"message": message,
	}

	g.ginCtx.JSON(statusCode, gin.H{
		"GinOut": []gin.H{errorItem},
	})
	g.ginCtx.Abort()
}
