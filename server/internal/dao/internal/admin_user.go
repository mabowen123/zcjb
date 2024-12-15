// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-12-14 13:54:21
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AdminUserDao is the data access object for the table admin_user.
type AdminUserDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of the current DAO.
	columns AdminUserColumns // columns contains all the column names of Table for convenient usage.
}

// AdminUserColumns defines and stores column names for the table admin_user.
type AdminUserColumns struct {
	Id        string // 主键
	Username  string // 用户名
	Password  string // 密码
	CreatedAt string // 创建时间戳
	UpdatedAt string // 更新时间戳
}

// adminUserColumns holds the columns for the table admin_user.
var adminUserColumns = AdminUserColumns{
	Id:        "id",
	Username:  "username",
	Password:  "password",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewAdminUserDao creates and returns a new DAO object for table data access.
func NewAdminUserDao() *AdminUserDao {
	return &AdminUserDao{
		group:   "zcjb",
		table:   "admin_user",
		columns: adminUserColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AdminUserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AdminUserDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AdminUserDao) Columns() AdminUserColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AdminUserDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AdminUserDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *AdminUserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
