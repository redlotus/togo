package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go/v4"
)

func JWTParserMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("Access-Token")
		ctx := r.Context()
		if tokenStr != "" {
			token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
				// TODO check validation of token also
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}

				return []byte(tokenSecret), nil
			})

			if err != nil {
				responseJSON(w, errInternalServer.Err(ErrorCodeTokenNotFound, fmt.Sprintf("failed to parse token: %v", err)), http.StatusInternalServerError)
				return
			}

			tokenClaims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				responseJSON(w, errInternalServer.Err(ErrorCodeTokenNotFound, fmt.Sprintf("failed to cast to target object: %v", err)), http.StatusInternalServerError)
				return
			}

			// FIXME parse to map[string]interface{}
			ctx = context.WithValue(ctx, "data", tokenClaims)
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
