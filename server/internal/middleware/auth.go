package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"server/internal/param/user"
	"server/internal/service"
	"server/utility"
)

func Auth(r *ghttp.Request) {
	token, err := jwt.ParseWithClaims(utility.GetToken(r), &user.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return utility.GetJwtKey(), nil
	})

	if err != nil || !token.Valid {
		r.Response.WriteStatusExit(http.StatusUnauthorized, nil)
	}

	userId := token.Claims.(*user.UserClaims).Id
	tokenStr := service.User().GetCacheTokenStr(r.Context(), userId)
	if utility.GetToken(r) != tokenStr {
		r.Response.WriteStatusExit(http.StatusUnauthorized, nil)
	}

	r.Middleware.Next()
}
