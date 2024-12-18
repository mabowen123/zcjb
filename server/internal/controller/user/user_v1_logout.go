package user

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"server/api/user/v1"
	"server/internal/service"
	"server/utility"
)

func (c *ControllerV1) Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error) {
	r := g.RequestFromCtx(ctx)
	tokenString := utility.GetToken(r)
	if tokenString == "" {
		return nil, nil
	}
	_, err = service.User().RemoveToken(ctx, tokenString)
	if err != nil {
		return nil, nil
	}

	return nil, nil
}
