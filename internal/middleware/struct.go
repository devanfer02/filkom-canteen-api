package middleware

import "github.com/devanfer02/filkom-canteen/internal/pkg/redis"

type Middleware struct {
	redis redis.RedisInterface	
}

func NewMiddleware(redis redis.RedisInterface) *Middleware {
	return &Middleware{redis: redis}
}