package v1

import "github.com/gogf/gf/v2/frame/g"

type CreateReq struct {
	g.Meta   `path:"/menu/create" tags:"菜单" method:"post" summary:"菜单创建"`
	Path     string `p:"path"  v:"required|max-length:20" dc:"菜单路由"`
	Name     string `p:"name"  v:"required|max-length:10" dc:"菜单名称"`
	Sort     uint   `p:"sort" dc:"排序"`
	ParentId uint   `p:"parent_id" dc:"菜单父节点id"`
}

type CreateRes struct {
	g.Meta `mime:"json"`
}
