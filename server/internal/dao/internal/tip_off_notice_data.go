// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-12-19 16:20:56
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TipOffNoticeDataDao is the data access object for the table tip_off_notice_data.
type TipOffNoticeDataDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of the current DAO.
	columns TipOffNoticeDataColumns // columns contains all the column names of Table for convenient usage.
}

// TipOffNoticeDataColumns defines and stores column names for the table tip_off_notice_data.
type TipOffNoticeDataColumns struct {
	Id            string //
	OrigId        string //
	Title         string //
	Content       string //
	DateTime      string //
	ShortTime     string //
	ShiJianChuo   string //
	CateId        string //
	CateName      string //
	Comments      string //
	LouZhu        string //
	LouZhuRegTime string //
	Url           string //
	PachongId     string //
	YuanUrl       string //
	IsNotice      string //
	CreatedAt     string // 创建时间戳
	UpdatedAt     string // 更新时间戳
}

// tipOffNoticeDataColumns holds the columns for the table tip_off_notice_data.
var tipOffNoticeDataColumns = TipOffNoticeDataColumns{
	Id:            "id",
	OrigId:        "orig_id",
	Title:         "title",
	Content:       "content",
	DateTime:      "date_time",
	ShortTime:     "short_time",
	ShiJianChuo:   "shi_jian_chuo",
	CateId:        "cate_id",
	CateName:      "cate_name",
	Comments:      "comments",
	LouZhu:        "lou_zhu",
	LouZhuRegTime: "lou_zhu_reg_time",
	Url:           "url",
	PachongId:     "pachong_id",
	YuanUrl:       "yuan_url",
	IsNotice:      "is_notice",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

// NewTipOffNoticeDataDao creates and returns a new DAO object for table data access.
func NewTipOffNoticeDataDao() *TipOffNoticeDataDao {
	return &TipOffNoticeDataDao{
		group:   "xianbao",
		table:   "tip_off_notice_data",
		columns: tipOffNoticeDataColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TipOffNoticeDataDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TipOffNoticeDataDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TipOffNoticeDataDao) Columns() TipOffNoticeDataColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TipOffNoticeDataDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TipOffNoticeDataDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *TipOffNoticeDataDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
