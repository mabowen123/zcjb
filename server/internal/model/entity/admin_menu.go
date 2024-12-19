// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-12-19 16:20:56
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminMenu is the golang structure for table admin_menu.
type AdminMenu struct {
	Id        uint        `json:"id"         orm:"id"         description:"主键"`     // 主键
	ParentId  uint        `json:"parent_id"  orm:"parent_id"  description:"父级菜单id"` // 父级菜单id
	Name      string      `json:"name"       orm:"name"       description:"菜单名称"`   // 菜单名称
	Path      string      `json:"path"       orm:"path"       description:"菜单路由"`   // 菜单路由
	Sort      uint        `json:"sort"       orm:"sort"       description:"排序"`     // 排序
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:"创建时间戳"`  // 创建时间戳
	UpdatedAt *gtime.Time `json:"updated_at" orm:"updated_at" description:"更新时间戳"`  // 更新时间戳
	DeletedAt *gtime.Time `json:"deleted_at" orm:"deleted_at" description:"删除时间戳"`  // 删除时间戳
}
