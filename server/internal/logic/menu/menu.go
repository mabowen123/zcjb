package menu

import (
	"context"
	"server/internal/consts/ierr"
	"server/internal/dao"
	"server/internal/model/entity"
	"server/internal/service"
)

type sMenu struct {
}

func NewMenu() *sMenu {
	return &sMenu{}
}
func init() {
	service.RegisterMenu(NewMenu())
}

func (*sMenu) Create(ctx context.Context, menu *entity.AdminMenu) (bool, error) {
	if menu.ParentId > 0 {
		isOk, err := dao.ExistById(ctx, menu.ParentId)
		if err != nil {
			return false, ierr.Common("查询父级异常")
		}

		if !isOk {
			return false, ierr.Common("父级不存在")
		}

		isOk, err = dao.ExistByPath(ctx, menu.Path)
		if err != nil {
			return false, ierr.Common("查询路由异常")
		}

		if isOk {
			return false, ierr.Common("路由已存在")
		}
	}

	_, err := dao.Create(ctx, menu)

	if err != nil {
		return false, ierr.Common("创建异常")
	}

	return true, nil
}
