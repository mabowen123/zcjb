package user

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"server/api/user/v1"
	"server/internal/param/user"
	"server/internal/service"
	"time"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	userinfo, err := service.User().GetUser(ctx, req.UserName, req.PassWord)
	if err != nil {
		return nil, err
	}

	userClaims := &user.UserClaims{
		Id:       userinfo.Id,
		Username: userinfo.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	res = new(v1.LoginRes)
	tokenStr, err := service.User().GetToken(ctx, userClaims)
	if err != nil {
		return nil, err
	}
	res.Token = tokenStr
	return res, nil
}
