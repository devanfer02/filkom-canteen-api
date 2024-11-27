package jwt

import (
	"fmt"

	j "github.com/golang-jwt/jwt/v5"

	"github.com/devanfer02/filkom-canteen/internal/infra/env"
	"github.com/devanfer02/filkom-canteen/internal/pkg/log"
)

type Claims struct {
	UserID string `json:"userId"`
	Role   string `json:"role"`
	j.RegisteredClaims
}

type Issuer struct {
	UserID string
	Issuer string
	Role   string
}

func ValidateToken(tokenReq string) (*Issuer, error) {
	var (
		err    error
		claims Claims
		token  *j.Token
		issuer *Issuer
	)

	token, err = j.ParseWithClaims(tokenReq, &claims, func(token *j.Token) (interface{}, error) {

		if _, ok := token.Method.(*j.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Method)
		}

		return []byte(env.AppEnv.JWTKey), nil
	})

	if err != nil {
		log.Info(log.LogInfo{
			"info_err": err.Error(),
		}, "[JWT][ValidateToken] failed to parse with claims")
		return nil, err
	}

	if !token.Valid {
		log.Info(log.LogInfo{
			"info_err": err.Error(),
		}, "[JWT][ValidateToken] token is invalid")

		return nil, err
	}

	issuer = &Issuer{
		UserID: claims.UserID,
		Issuer: claims.Issuer,
		Role:   claims.Role,
	}

	return issuer, nil
}
