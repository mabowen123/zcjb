// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-12-16 11:55:36
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminMenu is the golang structure of table admin_menu for DAO operations like Where/Data.
type AdminMenu struct {
	g.Meta    `orm:"table:admin_menu, do:true"`
	Id        interface{} // 主键
	ParentId  interface{} // 父级菜单id
	Name      interface{} // 菜单名称
	Path      interface{} // 菜单路由
	Sort      interface{} // 排序
	CreatedAt *gtime.Time // 创建时间戳
	UpdatedAt *gtime.Time // 更新时间戳
	DeletedAt *gtime.Time // 删除时间戳
}
