package ginlib

import (
	"github.com/devanfer02/filkom-canteen/domain"
	"github.com/gin-gonic/gin"
)

type Response struct {
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
	ctx.JSON(code, Response{
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

func SendAbortResponse(
	ctx *gin.Context,
	code int,
	status, message string,
	err error,	
) {
	SendResponse(ctx, code, status, message, nil, err)
	ctx.Abort()
}