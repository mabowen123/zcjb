// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-12-19 16:20:56
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ReptileConfigDao is the data access object for the table reptile_config.
type ReptileConfigDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of the current DAO.
	columns ReptileConfigColumns // columns contains all the column names of Table for convenient usage.
}

// ReptileConfigColumns defines and stores column names for the table reptile_config.
type ReptileConfigColumns struct {
	Id           string //
	Url          string // 请求地址
	IntervalTime string // 间隔时间(秒
	Type         string // 地址类型 1-线报
	NextTime     string // 下次执行时间
	Remark       string // 备注
	CreatedAt    string // 创建时间戳
	UpdatedAt    string // 更新时间戳
}

// reptileConfigColumns holds the columns for the table reptile_config.
var reptileConfigColumns = ReptileConfigColumns{
	Id:           "id",
	Url:          "url",
	IntervalTime: "interval_time",
	Type:         "type",
	NextTime:     "next_time",
	Remark:       "remark",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
}

// NewReptileConfigDao creates and returns a new DAO object for table data access.
func NewReptileConfigDao() *ReptileConfigDao {
	return &ReptileConfigDao{
		group:   "xianbao",
		table:   "reptile_config",
		columns: reptileConfigColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ReptileConfigDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ReptileConfigDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ReptileConfigDao) Columns() ReptileConfigColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ReptileConfigDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ReptileConfigDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *ReptileConfigDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}