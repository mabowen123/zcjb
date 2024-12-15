// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-12-15 21:43:40
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminUser is the golang structure of table admin_user for DAO operations like Where/Data.
type AdminUser struct {
	g.Meta    `orm:"table:admin_user, do:true"`
	Id        interface{} // 主键
	Username  interface{} // 用户名
	Password  interface{} // 密码
	CreatedAt *gtime.Time // 创建时间戳
	UpdatedAt *gtime.Time // 更新时间戳
}
