package middlewere

import (
	"fmt"
	"net/http"
	"strings"
	"web-chat/pkg/response"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func extractToken(r *http.Request) (string, error) {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1], nil
	}
	return "", fmt.Errorf("invalid token")
}

func verifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString, err := extractToken(r)
	if err != nil {
		return nil, err
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("config.AppConfig.JWTSecret"), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func tokenValid(r *http.Request) error {
	token, err := verifyToken(r)
	if err != nil {
		return err
	}
	if claim := token.Claims; claim.Valid() != nil {
		return err
	}
	return nil
}

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := tokenValid(ctx.Request)
		if err != nil {
			response.Unauthorized(ctx, err.Error())
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
