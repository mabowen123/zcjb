package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"server/internal/param/tip"
)

type ListReq struct {
	g.Meta  `path:"/tip/list" tags:"线报" method:"get" summary:"用户登录"`
	Page    string `p:"username"  v:"required"`
	PageNum string `p:"password"  v:"required"`
}

type ListRes struct {
	g.Meta `mime:"json"`
	Total  uint       `json:"total"`
	List   []tip.List `json:"list"`
}
