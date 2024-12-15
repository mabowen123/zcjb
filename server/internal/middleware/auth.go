package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"server/utility"
)

func Auth(r *ghttp.Request) {
	token, err := jwt.Parse(utility.GetToken(r), func(token *jwt.Token) (interface{}, error) {
		return utility.GetJwtKey(), nil
	})

	if err != nil || !token.Valid {
		r.Response.WriteStatus(http.StatusUnauthorized)
		r.ExitAll()
	}

	r.Middleware.Next()
}
