package v1

import "github.com/gogf/gf/v2/frame/g"

type LoginReq struct {
	g.Meta   `path:"/user/login" tags:"User" method:"post" summary:"用户登录"`
	UserName string `p:"username"  v:"required"`
	PassWord string `p:"password"  v:"required"`
}

type LoginRes struct {
	g.Meta `mime:"json"`
	Token  string `json:"token"`
}
