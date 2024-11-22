package middleware

// import (
// 	"errors"
// 	"strings"

// 	ginlib "github.com/devanfer02/filkom-canteen/internal/pkg/gin"
// 	"github.com/devanfer02/filkom-canteen/internal/pkg/jwt"
// 	"github.com/gin-gonic/gin"
// )

// func (m *Middleware) Authenticate() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		var (
// 			code = 400
// 			status = "fail"
// 			message = "failed to authenticate user"
// 			err error = nil 
// 		)

// 		bearer := ctx.GetHeader("Authorization")

// 		defer func() {
// 			if err != nil {
// 				ginlib.SendAbortResponse(ctx, code, status, message, err)
// 			}
// 		}()

// 		if bearer == "" {	
// 			err = errors.New("failed to get bearer token")
// 			return
// 		}
	
// 		splitted := strings.Split(bearer, " ")
	
// 		if len(splitted) < 2 {
// 			err = errors.New("failed to validate token")
// 			return
// 		}
	
// 		tokenString := splitted[1]
	
// 		id, user, role, err := jwt.ValidateToken(tokenString)
// 		if err != nil {
// 			err = errors.New("failed to validate token")
// 			return
// 		}
	
// 		if user != env.AppEnv.JwtUserRole && user != env.AppEnv.JwtAdminRole {
// 			response.SendErrResp(
// 				ctx,
// 				http.StatusUnauthorized,
// 				response.Fail,
// 				"failed to authenticate user",
// 				err,
// 			)
// 			return 
// 		}
	
// 		val, err := m.redis.Get(ctx.Request.Context(), tokenString)
	
// 		if err != nil {
// 			err = errors.New("token expired")
// 			return
// 		}
	
// 		if val != "" {
// 			err = errors.New("token expired")
// 			return
// 		}
	
// 		ctx.Set("id", id.String())
// 		ctx.Set("user", user)
// 		ctx.Set("role", role)
// 		ctx.Next()
// 	}
// }