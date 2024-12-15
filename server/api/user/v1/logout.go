package v1

import "github.com/gogf/gf/v2/frame/g"

type LogoutReq struct {
	g.Meta `path:"/user/logout" tags:"用户" method:"post" summary:"用户登出"`
}

type LogoutRes struct {
	g.Meta `mime:"json"`
}
