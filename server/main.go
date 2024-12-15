package main

import (
	_ "server/internal/logic"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"

	"server/internal/cmd"
	"server/utility"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	g.I18n().SetLanguage("zh-CN")
	utility.InitDB()
	utility.InitRedis()
	cmd.Main.Run(gctx.GetInitCtx())
}
