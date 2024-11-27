package middleware

import (
	"github.com/devanfer02/filkom-canteen/internal/app/repository"
	"github.com/devanfer02/filkom-canteen/internal/pkg/redis"
)

type Middleware struct {
	redis    redis.RedisInterface
	roleRepo repository.IRoleRepository
}

func NewMiddleware(redis redis.RedisInterface, roleRepo repository.IRoleRepository) *Middleware {
	return &Middleware{redis: redis, roleRepo: roleRepo}
}
