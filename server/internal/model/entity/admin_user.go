// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-12-19 16:20:56
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminUser is the golang structure for table admin_user.
type AdminUser struct {
	Id        uint        `json:"id"         orm:"id"         description:"主键"`    // 主键
	Username  string      `json:"username"   orm:"username"   description:"用户名"`   // 用户名
	Password  string      `json:"password"   orm:"password"   description:"密码"`    // 密码
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:"创建时间戳"` // 创建时间戳
	UpdatedAt *gtime.Time `json:"updated_at" orm:"updated_at" description:"更新时间戳"` // 更新时间戳
}
