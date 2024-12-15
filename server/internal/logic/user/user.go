package user

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/golang-jwt/jwt/v5"
	"server/internal/consts/ierr"
	"server/internal/dao"
	"server/internal/model/entity"
	"server/internal/param/user"
	"server/internal/service"
	"server/utility"
	"time"
)

type sUser struct {
}

func NewUser() *sUser {
	return &sUser{}
}
func init() {
	service.RegisterUser(NewUser())
}

func (s *sUser) GetUser(ctx context.Context, username string, password string) (*entity.AdminUser, error) {
	userinfo, err := dao.AdminUser.GetUserByUsernameAndPassword(ctx, username, password)

	if err != nil {
		g.Log().Errorf(ctx, "查询用户异常 username:%s  password:%s err:%s", username, password, err)
		return nil, ierr.Common("查询用户异常")
	}

	if userinfo == nil {
		return nil, ierr.Common("用户不存在")
	}

	return userinfo, nil
}

func (s *sUser) GetToken(ctx context.Context, user *user.UserClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, user)
	tokenStr, err := token.SignedString(utility.GetJwtKey())
	if err != nil {
		g.Log().Errorf(ctx, "生成token失败 user:%v err:%s", user, err)
		return "", ierr.Common("生成token失败")
	}
	dao.AdminUser.RSetExToken(ctx, user.Id, tokenStr, user.ExpiresAt.Unix()-time.Now().Unix())
	return tokenStr, nil
}

func (*sUser) RemoveToken(ctx context.Context, tokenStr string) (bool, error) {
	if tokenStr == "" {
		return true, nil
	}

	jwtKey := utility.GetJwtKey()
	tokenClaims, err := jwt.ParseWithClaims(tokenStr, &user.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		g.Log().Errorf(ctx, "解析token失败 token:%s err:%s", tokenStr, err)
		return true, ierr.Common("解析token失败")
	}

	if !tokenClaims.Valid {
		return true, nil
	}

	userClaims := tokenClaims.Claims.(*user.UserClaims)
	dao.AdminUser.RDelToken(ctx, userClaims.Id)
	return true, nil
}
