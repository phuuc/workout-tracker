package middleware

import (
	"github.com/finnpn/workout-tracker/pkg/errors"
	"github.com/finnpn/workout-tracker/pkg/ginout"
	"github.com/finnpn/workout-tracker/pkg/log"
	"github.com/gin-gonic/gin"
)

// ErrorMetadata ...
type ErrorMetadata struct {
	StatusCode   int
	ErrorCode    string
	ErrorMessage string
}

// Handleginout ...
func HandleErrors() gin.HandlerFunc {

	return func(c *gin.Context) {
		ginOut := ginout.NewGinOut(c)
		c.Next()
		lastError := c.Errors.Last()
		if lastError != nil {
			err := lastError.Err
			log.Info("middleware handle ginout ...")
			if lastError.Meta != nil {
				if errorMetadata, ok := lastError.Meta.(*ErrorMetadata); ok {
					ginOut.JSONError(errorMetadata.StatusCode, errorMetadata.ErrorCode, errorMetadata.ErrorMessage)
				}
			} else {
				ginOut.JSONError(errors.ErrInternalServerCode, errors.ErrInternalServer, err.Error())
			}
		}
	}
}
