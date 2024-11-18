package ginlib

import (
	"github.com/devanfer02/filkom-canteen/domain"
	"github.com/gin-gonic/gin"
)

type response struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Err     string      `json:"error,omitempty"`
}

func SendResponse(
	ctx *gin.Context,
	code int,
	status, message string,
	data interface{},
	err error,
) {
	ctx.JSON(code, response{
		Code:    code,
		Status:  status,
		Message: message,
		Data:    data,
		Err: func() string {
			if err == nil {
				return ""
			}

			if code, _ := domain.GetStatus(err); code == 500 {
				return "internal server error"
			}

			return err.Error()
		}(),
	})
}
