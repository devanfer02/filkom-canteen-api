package middleware

import (
	"errors"
	"strings"

	"github.com/devanfer02/filkom-canteen/internal/infra/env"
	ginlib "github.com/devanfer02/filkom-canteen/internal/pkg/gin"
	"github.com/gin-gonic/gin"
)

func APIKey() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			header  string
			code    int    = 400
			status  string = "fail"
			message        = "failed to authenticate request"
			err     error  = nil
		)

		defer func() {
			ginlib.SendAbortResponse(ctx, code, status, message, err)
		}()

		header = ctx.GetHeader("x-api-key")

		if header == "" {
			err = errors.New("invalid api key")
			return
		}

		split := strings.Split(header, " ")

		if len(split) < 2 {
			err = errors.New("invalid api key")
			return
		}

		if split[1] != env.AppEnv.ApiKey {
			err = errors.New("invalid api key")
		}

		ctx.Next()
	}
}
