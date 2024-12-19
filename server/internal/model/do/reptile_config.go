// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-12-19 16:20:56
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ReptileConfig is the golang structure of table reptile_config for DAO operations like Where/Data.
type ReptileConfig struct {
	g.Meta       `orm:"table:reptile_config, do:true"`
	Id           interface{} //
	Url          interface{} // 请求地址
	IntervalTime interface{} // 间隔时间(秒
	Type         interface{} // 地址类型 1-线报
	NextTime     interface{} // 下次执行时间
	Remark       interface{} // 备注
	CreatedAt    *gtime.Time // 创建时间戳
	UpdatedAt    *gtime.Time // 更新时间戳
}
