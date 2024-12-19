// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2024-12-19 16:20:56
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AdminMenuDao is the data access object for the table admin_menu.
type AdminMenuDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of the current DAO.
	columns AdminMenuColumns // columns contains all the column names of Table for convenient usage.
}

// AdminMenuColumns defines and stores column names for the table admin_menu.
type AdminMenuColumns struct {
	Id        string // 主键
	ParentId  string // 父级菜单id
	Name      string // 菜单名称
	Path      string // 菜单路由
	Sort      string // 排序
	CreatedAt string // 创建时间戳
	UpdatedAt string // 更新时间戳
	DeletedAt string // 删除时间戳
}

// adminMenuColumns holds the columns for the table admin_menu.
var adminMenuColumns = AdminMenuColumns{
	Id:        "id",
	ParentId:  "parent_id",
	Name:      "name",
	Path:      "path",
	Sort:      "sort",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewAdminMenuDao creates and returns a new DAO object for table data access.
func NewAdminMenuDao() *AdminMenuDao {
	return &AdminMenuDao{
		group:   "zcjb",
		table:   "admin_menu",
		columns: adminMenuColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AdminMenuDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AdminMenuDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AdminMenuDao) Columns() AdminMenuColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AdminMenuDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AdminMenuDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *AdminMenuDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
