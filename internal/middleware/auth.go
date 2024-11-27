package middleware

import (
	"errors"
	"strings"

	"github.com/devanfer02/filkom-canteen/domain"
	"github.com/devanfer02/filkom-canteen/internal/infra/env"
	ginlib "github.com/devanfer02/filkom-canteen/internal/pkg/gin"
	"github.com/devanfer02/filkom-canteen/internal/pkg/jwt"
	"github.com/devanfer02/filkom-canteen/internal/pkg/log"
	"github.com/gin-gonic/gin"
)

func (m *Middleware) Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			code = 400
			status = "fail"
			message = "failed to authenticate user"
			err error = nil 
		)

		bearer := ctx.GetHeader("Authorization")

		defer func() {
			if err != nil {
				ginlib.SendAbortResponse(ctx, code, status, message, err)
			}
		}()

		if bearer == "" {	
			err = errors.New("failed to get bearer token")
			return
		}
	
		splitted := strings.Split(bearer, " ")
	
		if len(splitted) < 2 {
			err = errors.New("failed to validate token")
			return
		}
	
		tokenString := splitted[1]
	
		issuer, err := jwt.ValidateToken(tokenString)
		if err != nil {
			err = errors.New("failed to validate token")
			return
		}

		if issuer.Issuer != env.AppEnv.JWTUserRole && issuer.Issuer != env.AppEnv.JWTAdminRole {
			err = errors.New("failed to validate token")
			return 
		}

	
		val, err := m.redis.Get(ctx.Request.Context(), tokenString)
	
		if err != nil {
			err = errors.New("token expired")
			return
		}
	
		if val != "" {
			err = errors.New("token expired")
			return
		}
	
		ctx.Set("id", issuer.UserID)
		ctx.Set("user", issuer.Issuer)
		ctx.Set("role", issuer.Role)
		ctx.Next()
	}
}

func (m *Middleware) AuthorizeAdmin(roleName string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			err error 
			code = 401
			status = "fail"
			message = "unauthorized"
			role *domain.Role
		)

		role, err = m.roleRepo.FetchOne(roleName)

		defer func() {
			if err != nil {
				ginlib.SendAbortResponse(ctx, code, status, message, err)
			}
		}()

		if err != nil {
			err = errors.New("failed to authorize")
			return 
		}

		if role.ID != ctx.GetString("role") {
			log.Info(log.LogInfo{
				"role_db": role.ID,
				"role_ctx": ctx.GetString("role"),
			}, "LOGGED")
			err = errors.New("unauthorized")
			return 
		}

		ctx.Next()
	}
}