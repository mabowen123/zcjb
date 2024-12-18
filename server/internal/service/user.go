// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"server/internal/model/entity"
	"server/internal/param/user"
)

type (
	IUser interface {
		GetUser(ctx context.Context, username string, password string) (*entity.AdminUser, error)
		GetToken(ctx context.Context, user *user.UserClaims) (string, error)
		RemoveToken(ctx context.Context, tokenStr string) (bool, error)
		GetCacheTokenStr(ctx context.Context, id uint) string
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
