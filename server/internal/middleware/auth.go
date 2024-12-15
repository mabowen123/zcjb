package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/golang-jwt/jwt/v5"
	"server/utility"
)

func Auth(r *ghttp.Request) {
	var (
		jwtKey      = utility.GetJwtKey()
		tokenString = r.Header.Get("Auth")
	)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	//fmt.Println(token.Claims)
	if err != nil || !token.Valid {
		//r.Response.WriteStatus(http.StatusForbidden)
		//r.Exit()
	}

	r.Middleware.Next()
}
