package utility

import "github.com/gogf/gf/v2/net/ghttp"

func GetToken(r *ghttp.Request) string {
	return r.Header.Get("auth")
}
