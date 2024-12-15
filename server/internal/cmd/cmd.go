package cmd

import (
	"context"
	"fmt"
	"server/internal/controller/user"
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
			s.Group("admin", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Middleware(ghttp.MiddlewareCORS)
				group.Bind(
					user.NewV1(),
				)
			})

			s.Run()
			fmt.Println(ctx)
			return nil
		},
	}
)
