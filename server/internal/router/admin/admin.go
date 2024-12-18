package admin

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"server/internal/controller/user"
	"server/internal/middleware"
)

func RegisterRouter(group *ghttp.RouterGroup) {
	group.Group("/admin", func(group *ghttp.RouterGroup) {
		group.Bind(
			user.NewV1(),
		)

		group.Middleware(middleware.Auth)

	})
}
