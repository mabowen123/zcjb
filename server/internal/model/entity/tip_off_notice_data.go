// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-12-19 16:20:56
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TipOffNoticeData is the golang structure for table tip_off_notice_data.
type TipOffNoticeData struct {
	Id            uint        `json:"id"               orm:"id"               description:""`      //
	OrigId        int64       `json:"orig_id"          orm:"orig_id"          description:""`      //
	Title         string      `json:"title"            orm:"title"            description:""`      //
	Content       string      `json:"content"          orm:"content"          description:""`      //
	DateTime      string      `json:"date_time"        orm:"date_time"        description:""`      //
	ShortTime     string      `json:"short_time"       orm:"short_time"       description:""`      //
	ShiJianChuo   uint64      `json:"shi_jian_chuo"    orm:"shi_jian_chuo"    description:""`      //
	CateId        string      `json:"cate_id"          orm:"cate_id"          description:""`      //
	CateName      string      `json:"cate_name"        orm:"cate_name"        description:""`      //
	Comments      uint        `json:"comments"         orm:"comments"         description:""`      //
	LouZhu        string      `json:"lou_zhu"          orm:"lou_zhu"          description:""`      //
	LouZhuRegTime string      `json:"lou_zhu_reg_time" orm:"lou_zhu_reg_time" description:""`      //
	Url           string      `json:"url"              orm:"url"              description:""`      //
	PachongId     int         `json:"pachong_id"       orm:"pachong_id"       description:""`      //
	YuanUrl       string      `json:"yuan_url"         orm:"yuan_url"         description:""`      //
	IsNotice      uint        `json:"is_notice"        orm:"is_notice"        description:""`      //
	CreatedAt     *gtime.Time `json:"created_at"       orm:"created_at"       description:"创建时间戳"` // 创建时间戳
	UpdatedAt     *gtime.Time `json:"updated_at"       orm:"updated_at"       description:"更新时间戳"` // 更新时间戳
}
