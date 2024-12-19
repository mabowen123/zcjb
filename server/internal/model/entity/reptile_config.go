// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-12-19 16:20:56
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ReptileConfig is the golang structure for table reptile_config.
type ReptileConfig struct {
	Id           uint        `json:"id"            orm:"id"            description:""`          //
	Url          string      `json:"url"           orm:"url"           description:"请求地址"`      // 请求地址
	IntervalTime uint        `json:"interval_time" orm:"interval_time" description:"间隔时间(秒"`    // 间隔时间(秒
	Type         uint        `json:"type"          orm:"type"          description:"地址类型 1-线报"` // 地址类型 1-线报
	NextTime     uint64      `json:"next_time"     orm:"next_time"     description:"下次执行时间"`    // 下次执行时间
	Remark       string      `json:"remark"        orm:"remark"        description:"备注"`        // 备注
	CreatedAt    *gtime.Time `json:"created_at"    orm:"created_at"    description:"创建时间戳"`     // 创建时间戳
	UpdatedAt    *gtime.Time `json:"updated_at"    orm:"updated_at"    description:"更新时间戳"`     // 更新时间戳
}
