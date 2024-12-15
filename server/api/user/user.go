// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package user

import (
	"context"

	"server/api/user/v1"
)

type IUserV1 interface {
	Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error)
	Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error)
}
