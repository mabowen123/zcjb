// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-12-19 16:20:56
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TipOffNoticeData is the golang structure of table tip_off_notice_data for DAO operations like Where/Data.
type TipOffNoticeData struct {
	g.Meta        `orm:"table:tip_off_notice_data, do:true"`
	Id            interface{} //
	OrigId        interface{} //
	Title         interface{} //
	Content       interface{} //
	DateTime      interface{} //
	ShortTime     interface{} //
	ShiJianChuo   interface{} //
	CateId        interface{} //
	CateName      interface{} //
	Comments      interface{} //
	LouZhu        interface{} //
	LouZhuRegTime interface{} //
	Url           interface{} //
	PachongId     interface{} //
	YuanUrl       interface{} //
	IsNotice      interface{} //
	CreatedAt     *gtime.Time // 创建时间戳
	UpdatedAt     *gtime.Time // 更新时间戳
}
