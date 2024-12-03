package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	middleware "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

func (m *Middleware) RateLimiter(limit int64) gin.HandlerFunc {
	var (
		rate = limiter.Rate{
			Period: 1 * time.Hour,
			Limit: limit,
		}

		store = memory.NewStore()

		instance = limiter.New(store, rate)
	)

	return middleware.NewMiddleware(instance)
}
