package cmd

import (
	"context"
	"server/internal/middleware"
	"server/internal/router/admin"
	"server/utility"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  utility.GetProjectName(),
		Usage: utility.GetProjectName(),
		Func: func(ctx context.Context, _ *gcmd.Parser) (err error) {
			s := g.Server()
			s.Use(middleware.Response, ghttp.MiddlewareCORS)
			s.Group("/", func(group *ghttp.RouterGroup) {
				admin.RegisterRouter(group)
			})

			s.Run()
			return nil
		},
	}
)
