package menu

import (
	"context"
	"github.com/gogf/gf/v2/text/gstr"
	"server/api/menu/v1"
	"server/internal/model/entity"
	"server/internal/service"
	"strings"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	path := gstr.Replace(gstr.Replace(req.Path, " ", ""), "//", "/")
	if req.ParentId == 0 {
		//父级菜单不跳转
		path = "/"
	} else if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	name := gstr.Replace(req.Name, " ", "")

	_, err = service.Menu().Create(ctx, &entity.AdminMenu{
		ParentId: req.ParentId,
		Name:     name,
		Path:     path,
		Sort:     req.Sort,
	})

	return
}
